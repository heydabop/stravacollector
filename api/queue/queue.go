package queue

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/Emyrk/strava/strava/stravalimit"

	"golang.org/x/oauth2"

	"github.com/Emyrk/strava/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	exp "github.com/vgarvardt/backoff"
	"github.com/vgarvardt/gue/v5"
	"github.com/vgarvardt/gue/v5/adapter/pgxv5"
	zadapter "github.com/vgarvardt/gue/v5/adapter/zerolog"
)

const (
	fetchActivityJob    = "fetch_activity"
	updateActivityField = "update_activity"
	deleteActivityJob   = "delete_activity"

	stravaFetchQueue          = "queue_strava_fetch"
	stravaUpdateActivityQueue = "queue_strava_update_activity"
)

var (
	rateLimitJobFail = errors.New("hitting strava rate limit, failing job to try later")
)

type Options struct {
	DBURL    string
	Logger   zerolog.Logger
	DB       database.Store
	OAuthCfg *oauth2.Config
	Registry *prometheus.Registry
}

// Manager will handle all queue related operations and jobs
type Manager struct {
	Client *gue.Client

	// Pool is used for the queuing library
	pool *pgxpool.Pool

	// DB is used by jobs
	DB database.Store

	Logger   zerolog.Logger
	OAuthCfg *oauth2.Config

	cancel              context.CancelFunc
	stravaLimitDebounce atomic.Pointer[time.Time]

	// Metrics!
	backloadHistogram        *prometheus.HistogramVec
	backgroundJobHistogram   *prometheus.HistogramVec
	backloadActivitiesLoaded prometheus.Counter
}

func New(ctx context.Context, opts Options) (*Manager, error) {
	cfg, err := pgxpool.ParseConfig(opts.DBURL)
	if err != nil {
		return nil, fmt.Errorf("parse postgres db url: %w", err)
	}
	// Small number of conns
	cfg.MaxConns = 2

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("new pool: %w", err)
	}

	poolAdapter := pgxv5.NewConnPool(pool)
	cli, err := gue.NewClient(poolAdapter,
		gue.WithClientLogger(zadapter.New(opts.Logger)),
		gue.WithClientBackoff(gue.NewExponentialBackoff(exp.Config{
			BaseDelay:  time.Second * 5,
			Multiplier: 1.6,
			Jitter:     0.2,
			MaxDelay:   time.Minute * 15,
		})),
	)
	if err != nil {
		return nil, fmt.Errorf("new client: %w", err)
	}
	if opts.Registry == nil {
		opts.Registry = prometheus.NewRegistry()
	}

	factory := promauto.With(opts.Registry)
	return &Manager{
		Client:   cli,
		pool:     pool,
		DB:       opts.DB,
		OAuthCfg: opts.OAuthCfg,
		Logger:   opts.Logger,

		// Metrics!
		backloadHistogram: factory.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "strava",
			Subsystem: "manager",
			Name:      "backload_athlete_seconds",
			Help:      "Each time we backload an athlete",
			// A 1s backload is fine imo
			Buckets: []float64{0.001, 0.005, 0.01, 0.25, 0.5, 1, 2, 5},
		}, []string{"success"}),
		backloadActivitiesLoaded: factory.NewCounter(prometheus.CounterOpts{
			Namespace: "strava",
			Subsystem: "manager",
			Name:      "backload_activities_enqueued_total",
			Help:      "The total number of activities enqueued to be fetched",
		}),
		backgroundJobHistogram: factory.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "strava",
			Subsystem: "manager",
			Name:      "background_job_seconds",
			Help:      "Each time we run a background job",
			Buckets:   []float64{0.001, 0.005, 0.01, 0.25, 0.5, 1, 2, 5},
		}, []string{"type", "success"}),
	}, nil
}

func (m *Manager) failedJobHook() func(ctx context.Context, j *gue.Job, err error) {
	return func(ctx context.Context, j *gue.Job, err error) {
		// TODO: If this is a strava too many requests, we need to sleep.
		if err != nil {
			if errors.Is(err, rateLimitJobFail) {
				return
			}
			m.Logger.Error().
				Err(err).
				Str("job_id", j.ID.String()).
				Str("job", j.Type).
				Str("queue", j.Queue).
				Int32("err_count", j.ErrorCount).
				Str("last_error", j.LastError.String).
				Msg("job failed")
		}
	}
}

func (m *Manager) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	m.cancel = cancel

	// worker for strava fetch queue
	workers, err := m.newWorkers([]string{stravaFetchQueue, stravaUpdateActivityQueue})
	if err != nil {
		return fmt.Errorf("new workers: %w", err)
	}

	for _, w := range workers {
		w := w
		// TODO: Errogroup these guys
		go func(w *gue.Worker) {
			err := w.Run(ctx)
			if err != nil {
				m.Logger.Error().Err(err).Msg("worker error")
			}
			cancel()
		}(w)
	}

	// Run backloading!
	go func() {
		// TODO: Make this able to scale horizontally
		m.BackLoadAthleteRoutine(ctx)
	}()

	go func() {
		// TODO: Make this able to scale horizontally
		m.BackLoadRouteSegments(ctx)
	}()

	return nil
}

func (m *Manager) newWorkers(queues []string, opts ...gue.WorkerOption) ([]*gue.Worker, error) {
	var workers []*gue.Worker
	for _, q := range queues {
		qOpts := make([]gue.WorkerOption, len(opts))
		copy(qOpts, opts)

		worker, err := m.newWorker(q, qOpts...)
		if err != nil {
			return nil, fmt.Errorf("new worker %s: %w", q, err)
		}
		workers = append(workers, worker)
	}
	return workers, nil
}

func (m *Manager) newWorker(queue string, opts ...gue.WorkerOption) (*gue.Worker, error) {
	opts = append(opts,
		gue.WithWorkerQueue(queue),
		gue.WithWorkerHooksJobDone(m.failedJobHook()),
	)
	// All workers share the workmap
	worker, err := gue.NewWorker(m.Client, m.workMap(),
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("new worker: %w", err)
	}

	return worker, nil
}

func (m *Manager) workMap() gue.WorkMap {
	return gue.WorkMap{
		"online": func(ctx context.Context, j *gue.Job) error {
			m.Logger.Info().Msg("worker online")
			return nil
		},
		fetchActivityJob:    m.instrumentJob(m.fetchActivity),
		updateActivityField: m.instrumentJob(m.updateActivity),
		deleteActivityJob:   m.instrumentJob(m.deleteActivity),
	}
}

func (m *Manager) instrumentJob(runJob func(ctx context.Context, j *gue.Job) error) func(ctx context.Context, j *gue.Job) error {
	return func(ctx context.Context, j *gue.Job) error {
		start := time.Now()
		err := runJob(ctx, j)
		m.backgroundJobHistogram.
			WithLabelValues(j.Type, strconv.FormatBool(err == nil)).
			Observe(time.Since(start).Seconds())
		return err
	}
}

func (m *Manager) jobStravaCheck(j *gue.Job, calls int64) error {
	logger := jobLogFields(m.Logger, j)
	iBuf, dBuf := int64(100), int64(500)
	if stravalimit.NextDailyReset(time.Now()) < time.Hour*3 {
		iBuf, dBuf = int64(50), int64(100)
	}

	ok, limitLogger := stravalimit.CanLogger(1, iBuf, dBuf, logger)
	if !ok {
		last := m.stravaLimitDebounce.Load()
		now := time.Now()
		if last == nil || now.Sub(*last) > time.Minute*5 {
			limitLogger.Error().
				Msg("hitting strava rate limit, job going to fail and try again later")
			m.stravaLimitDebounce.Store(&now)
		}
		return rateLimitJobFail
	}
	return nil
}

func jobLogFields(logger zerolog.Logger, j *gue.Job) zerolog.Logger {
	return logger.With().
		Str("job_id", j.ID.String()).
		Str("job", j.Type).
		Str("queue", j.Queue).
		Int32("err_count", j.ErrorCount).
		Str("last_error", j.LastError.String).
		Logger()
}

func (m *Manager) Close() {
	m.cancel()
	m.pool.Close()
}
