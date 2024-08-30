// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"context"
)

type sqlcQuerier interface {
	AllCompetitiveRoutes(ctx context.Context) ([]CompetitiveRoute, error)
	AthleteHugelActivites(ctx context.Context, athleteID int64) ([]AthleteHugelActivitesRow, error)
	AthleteSyncedActivities(ctx context.Context, arg AthleteSyncedActivitiesParams) ([]AthleteSyncedActivitiesRow, error)
	// BestRouteEfforts returns all activities that have efforts on all the provided segments.
	// The returned activities include the best effort for each segment.
	// This isn't used in the app, but is the foundation for the hugel view.
	BestRouteEfforts(ctx context.Context, expectedSegments []int64) ([]BestRouteEffortsRow, error)
	DeleteActivity(ctx context.Context, id int64) (ActivitySummary, error)
	DeleteAthleteLogin(ctx context.Context, athleteID int64) error
	GetActivityDetail(ctx context.Context, id int64) (ActivityDetail, error)
	GetActivitySummary(ctx context.Context, id int64) (ActivitySummary, error)
	GetAthlete(ctx context.Context, athleteID int64) (Athlete, error)
	GetAthleteFull(ctx context.Context, athleteID int64) (GetAthleteFullRow, error)
	GetAthleteLoad(ctx context.Context, athleteID int64) (AthleteLoad, error)
	GetAthleteLoadDetailed(ctx context.Context, athleteID int64) (GetAthleteLoadDetailedRow, error)
	GetAthleteLogin(ctx context.Context, athleteID int64) (AthleteLogin, error)
	GetAthleteLoginFull(ctx context.Context, athleteID int64) (GetAthleteLoginFullRow, error)
	GetAthleteNeedsLoad(ctx context.Context) ([]GetAthleteNeedsLoadRow, error)
	GetCompetitiveRoute(ctx context.Context, routeName string) (GetCompetitiveRouteRow, error)
	// For authenticated users
	GetPersonalSegments(ctx context.Context, arg GetPersonalSegmentsParams) ([]GetPersonalSegmentsRow, error)
	GetSegments(ctx context.Context, segmentIds []int64) ([]GetSegmentsRow, error)
	HugelLeaderboard(ctx context.Context, arg HugelLeaderboardParams) ([]HugelLeaderboardRow, error)
	InsertFailedJob(ctx context.Context, rawJson string) (FailedJob, error)
	InsertWebhookDump(ctx context.Context, rawJson string) (WebhookDump, error)
	LoadedSegments(ctx context.Context) ([]LoadedSegmentsRow, error)
	MissingSegments(ctx context.Context, activitiesID int64) ([]string, error)
	NeedsARefresh(ctx context.Context) ([]NeedsARefreshRow, error)
	RefreshHugelActivities(ctx context.Context) error
	StarSegments(ctx context.Context, arg StarSegmentsParams) error
	SuperHugelLeaderboard(ctx context.Context, athleteID interface{}) ([]SuperHugelLeaderboardRow, error)
	TotalActivityDetailsCount(ctx context.Context) (int64, error)
	TotalJobCount(ctx context.Context) (int64, error)
	TotalRideActivitySummariesCount(ctx context.Context) (int64, error)
	UpdateActivityName(ctx context.Context, arg UpdateActivityNameParams) error
	UpsertActivityDetail(ctx context.Context, arg UpsertActivityDetailParams) (ActivityDetail, error)
	UpsertActivitySummary(ctx context.Context, arg UpsertActivitySummaryParams) (ActivitySummary, error)
	UpsertAthlete(ctx context.Context, arg UpsertAthleteParams) (Athlete, error)
	UpsertAthleteLoad(ctx context.Context, arg UpsertAthleteLoadParams) (AthleteLoad, error)
	UpsertAthleteLogin(ctx context.Context, arg UpsertAthleteLoginParams) (AthleteLogin, error)
	UpsertMapData(ctx context.Context, arg UpsertMapDataParams) (Map, error)
	UpsertSegment(ctx context.Context, arg UpsertSegmentParams) (Segment, error)
	UpsertSegmentEffort(ctx context.Context, arg UpsertSegmentEffortParams) (SegmentEffort, error)
	test(ctx context.Context) ([]testRow, error)
}

var _ sqlcQuerier = (*sqlQuerier)(nil)
