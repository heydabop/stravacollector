// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// The source of the activity fetching.
type ActivityDetailSource string

const (
	ActivityDetailSourceWebhook            ActivityDetailSource = "webhook"
	ActivityDetailSourceBackload           ActivityDetailSource = "backload"
	ActivityDetailSourceRequested          ActivityDetailSource = "requested"
	ActivityDetailSourceManual             ActivityDetailSource = "manual"
	ActivityDetailSourceUnknown            ActivityDetailSource = "unknown"
	ActivityDetailSourceZeroSegmentRefetch ActivityDetailSource = "zero_segment_refetch"
)

func (e *ActivityDetailSource) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ActivityDetailSource(s)
	case string:
		*e = ActivityDetailSource(s)
	default:
		return fmt.Errorf("unsupported scan type for ActivityDetailSource: %T", src)
	}
	return nil
}

type NullActivityDetailSource struct {
	ActivityDetailSource ActivityDetailSource `json:"activity_detail_source"`
	Valid                bool                 `json:"valid"` // Valid is true if ActivityDetailSource is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullActivityDetailSource) Scan(value interface{}) error {
	if value == nil {
		ns.ActivityDetailSource, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ActivityDetailSource.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullActivityDetailSource) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ActivityDetailSource), nil
}

func (e ActivityDetailSource) Valid() bool {
	switch e {
	case ActivityDetailSourceWebhook,
		ActivityDetailSourceBackload,
		ActivityDetailSourceRequested,
		ActivityDetailSourceManual,
		ActivityDetailSourceUnknown,
		ActivityDetailSourceZeroSegmentRefetch:
		return true
	}
	return false
}

func AllActivityDetailSourceValues() []ActivityDetailSource {
	return []ActivityDetailSource{
		ActivityDetailSourceWebhook,
		ActivityDetailSourceBackload,
		ActivityDetailSourceRequested,
		ActivityDetailSourceManual,
		ActivityDetailSourceUnknown,
		ActivityDetailSourceZeroSegmentRefetch,
	}
}

type ActivityDetail struct {
	ID                       int64     `db:"id" json:"id"`
	AthleteID                int64     `db:"athlete_id" json:"athlete_id"`
	StartLatlng              []float64 `db:"start_latlng" json:"start_latlng"`
	EndLatlng                []float64 `db:"end_latlng" json:"end_latlng"`
	FromAcceptedTag          bool      `db:"from_accepted_tag" json:"from_accepted_tag"`
	AverageCadence           float64   `db:"average_cadence" json:"average_cadence"`
	AverageTemp              float64   `db:"average_temp" json:"average_temp"`
	AverageWatts             float64   `db:"average_watts" json:"average_watts"`
	WeightedAverageWatts     float64   `db:"weighted_average_watts" json:"weighted_average_watts"`
	Kilojoules               float64   `db:"kilojoules" json:"kilojoules"`
	MaxWatts                 float64   `db:"max_watts" json:"max_watts"`
	ElevHigh                 float64   `db:"elev_high" json:"elev_high"`
	ElevLow                  float64   `db:"elev_low" json:"elev_low"`
	SufferScore              int32     `db:"suffer_score" json:"suffer_score"`
	Calories                 float64   `db:"calories" json:"calories"`
	EmbedToken               string    `db:"embed_token" json:"embed_token"`
	SegmentLeaderboardOptOut bool      `db:"segment_leaderboard_opt_out" json:"segment_leaderboard_opt_out"`
	LeaderboardOptOut        bool      `db:"leaderboard_opt_out" json:"leaderboard_opt_out"`
	NumSegmentEfforts        int32     `db:"num_segment_efforts" json:"num_segment_efforts"`
	// Owner of the activity has premium account at the time of the fetch.
	PremiumFetch bool `db:"premium_fetch" json:"premium_fetch"`
	// The time at which the activity was last updated by the collector
	UpdatedAt time.Time            `db:"updated_at" json:"updated_at"`
	MapID     string               `db:"map_id" json:"map_id"`
	Source    ActivityDetailSource `db:"source" json:"source"`
}

// Activity is missing many detailed fields
type ActivitySummary struct {
	ID                 int64     `db:"id" json:"id"`
	AthleteID          int64     `db:"athlete_id" json:"athlete_id"`
	UploadID           int64     `db:"upload_id" json:"upload_id"`
	ExternalID         string    `db:"external_id" json:"external_id"`
	Name               string    `db:"name" json:"name"`
	Distance           float64   `db:"distance" json:"distance"`
	MovingTime         float64   `db:"moving_time" json:"moving_time"`
	ElapsedTime        float64   `db:"elapsed_time" json:"elapsed_time"`
	TotalElevationGain float64   `db:"total_elevation_gain" json:"total_elevation_gain"`
	ActivityType       string    `db:"activity_type" json:"activity_type"`
	SportType          string    `db:"sport_type" json:"sport_type"`
	WorkoutType        int32     `db:"workout_type" json:"workout_type"`
	StartDate          time.Time `db:"start_date" json:"start_date"`
	StartDateLocal     time.Time `db:"start_date_local" json:"start_date_local"`
	Timezone           string    `db:"timezone" json:"timezone"`
	UtcOffset          float64   `db:"utc_offset" json:"utc_offset"`
	AchievementCount   int32     `db:"achievement_count" json:"achievement_count"`
	KudosCount         int32     `db:"kudos_count" json:"kudos_count"`
	CommentCount       int32     `db:"comment_count" json:"comment_count"`
	AthleteCount       int32     `db:"athlete_count" json:"athlete_count"`
	PhotoCount         int32     `db:"photo_count" json:"photo_count"`
	MapID              string    `db:"map_id" json:"map_id"`
	Trainer            bool      `db:"trainer" json:"trainer"`
	Commute            bool      `db:"commute" json:"commute"`
	Manual             bool      `db:"manual" json:"manual"`
	Private            bool      `db:"private" json:"private"`
	Flagged            bool      `db:"flagged" json:"flagged"`
	GearID             string    `db:"gear_id" json:"gear_id"`
	AverageSpeed       float64   `db:"average_speed" json:"average_speed"`
	MaxSpeed           float64   `db:"max_speed" json:"max_speed"`
	DeviceWatts        bool      `db:"device_watts" json:"device_watts"`
	HasHeartrate       bool      `db:"has_heartrate" json:"has_heartrate"`
	PrCount            int32     `db:"pr_count" json:"pr_count"`
	TotalPhotoCount    int32     `db:"total_photo_count" json:"total_photo_count"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
	AverageHeartrate   float64   `db:"average_heartrate" json:"average_heartrate"`
	MaxHeartrate       float64   `db:"max_heartrate" json:"max_heartrate"`
	DownloadCount      int32     `db:"download_count" json:"download_count"`
}

type Athlete struct {
	ID          int64  `db:"id" json:"id"`
	Summit      bool   `db:"summit" json:"summit"`
	Username    string `db:"username" json:"username"`
	Firstname   string `db:"firstname" json:"firstname"`
	Lastname    string `db:"lastname" json:"lastname"`
	Sex         string `db:"sex" json:"sex"`
	City        string `db:"city" json:"city"`
	State       string `db:"state" json:"state"`
	Country     string `db:"country" json:"country"`
	FollowCount int32  `db:"follow_count" json:"follow_count"`
	FriendCount int32  `db:"friend_count" json:"friend_count"`
	// feet or meters
	MeasurementPreference string          `db:"measurement_preference" json:"measurement_preference"`
	Ftp                   float64         `db:"ftp" json:"ftp"`
	Weight                float64         `db:"weight" json:"weight"`
	Clubs                 json.RawMessage `db:"clubs" json:"clubs"`
	CreatedAt             time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt             time.Time       `db:"updated_at" json:"updated_at"`
	FetchedAt             time.Time       `db:"fetched_at" json:"fetched_at"`
	ProfilePicLink        string          `db:"profile_pic_link" json:"profile_pic_link"`
	ProfilePicLinkMedium  string          `db:"profile_pic_link_medium" json:"profile_pic_link_medium"`
}

type AthleteHugelCount struct {
	AthleteID int64 `db:"athlete_id" json:"athlete_id"`
	Count     int64 `db:"count" json:"count"`
}

type AthleteHugelCount2023 struct {
	AthleteID int64 `db:"athlete_id" json:"athlete_id"`
	Count     int64 `db:"count" json:"count"`
}

// Tracks loading athlete activities. Must be an authenticated athlete.
type AthleteLoad struct {
	AthleteID int64 `db:"athlete_id" json:"athlete_id"`
	// Timestamp start of the last activity loaded. Future ones are not loaded.
	LastBackloadActivityStart time.Time `db:"last_backload_activity_start" json:"last_backload_activity_start"`
	// Timestamp of the last time the athlete was attempted to be loaded.
	LastLoadAttempt time.Time `db:"last_load_attempt" json:"last_load_attempt"`
	// True if the last load was incomplete and needs more work to catch up.
	LastLoadIncomplete         bool   `db:"last_load_incomplete" json:"last_load_incomplete"`
	LastLoadError              string `db:"last_load_error" json:"last_load_error"`
	ActivitesLoadedLastAttempt int32  `db:"activites_loaded_last_attempt" json:"activites_loaded_last_attempt"`
	// The earliest activity found for the athlete
	EarliestActivity time.Time `db:"earliest_activity" json:"earliest_activity"`
	// Loading backwards is done
	EarliestActivityDone bool      `db:"earliest_activity_done" json:"earliest_activity_done"`
	EarliestActivityID   int64     `db:"earliest_activity_id" json:"earliest_activity_id"`
	NextLoadNotBefore    time.Time `db:"next_load_not_before" json:"next_load_not_before"`
	CreatedAt            time.Time `db:"created_at" json:"created_at"`
}

type AthleteLogin struct {
	AthleteID int64 `db:"athlete_id" json:"athlete_id"`
	Summit    bool  `db:"summit" json:"summit"`
	// Oauth app client ID
	ProviderID        string    `db:"provider_id" json:"provider_id"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
	OauthAccessToken  string    `db:"oauth_access_token" json:"oauth_access_token"`
	OauthRefreshToken string    `db:"oauth_refresh_token" json:"oauth_refresh_token"`
	OauthExpiry       time.Time `db:"oauth_expiry" json:"oauth_expiry"`
	OauthTokenType    string    `db:"oauth_token_type" json:"oauth_token_type"`
	ID                uuid.UUID `db:"id" json:"id"`
}

type CompetitiveRoute struct {
	Name        string  `db:"name" json:"name"`
	DisplayName string  `db:"display_name" json:"display_name"`
	Description string  `db:"description" json:"description"`
	Segments    []int64 `db:"segments" json:"segments"`
}

// A table to store failed job information for potential debugging.
type FailedJob struct {
	// Some random uuid
	ID         uuid.UUID `db:"id" json:"id"`
	RecordedAt time.Time `db:"recorded_at" json:"recorded_at"`
	// Some text. Probably a JSON string.
	Raw string `db:"raw" json:"raw"`
}

type GueJob struct {
	JobID      string         `db:"job_id" json:"job_id"`
	Priority   int16          `db:"priority" json:"priority"`
	RunAt      time.Time      `db:"run_at" json:"run_at"`
	JobType    string         `db:"job_type" json:"job_type"`
	Args       []byte         `db:"args" json:"args"`
	ErrorCount int32          `db:"error_count" json:"error_count"`
	LastError  sql.NullString `db:"last_error" json:"last_error"`
	Queue      string         `db:"queue" json:"queue"`
	CreatedAt  time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at" json:"updated_at"`
}

type HugelActivities2023 struct {
	ActivityID       int64           `db:"activity_id" json:"activity_id"`
	AthleteID        int64           `db:"athlete_id" json:"athlete_id"`
	SegmentIds       interface{}     `db:"segment_ids" json:"segment_ids"`
	TotalTimeSeconds int64           `db:"total_time_seconds" json:"total_time_seconds"`
	Efforts          json.RawMessage `db:"efforts" json:"efforts"`
}

type HugelActivity struct {
	ActivityID       int64               `db:"activity_id" json:"activity_id"`
	AthleteID        int64               `db:"athlete_id" json:"athlete_id"`
	SegmentIds       interface{}         `db:"segment_ids" json:"segment_ids"`
	TotalTimeSeconds int64               `db:"total_time_seconds" json:"total_time_seconds"`
	Efforts          HugelSegmentEfforts `db:"efforts" json:"efforts"`
}

type LiteHugelActivity struct {
	ActivityID       int64           `db:"activity_id" json:"activity_id"`
	AthleteID        int64           `db:"athlete_id" json:"athlete_id"`
	SegmentIds       interface{}     `db:"segment_ids" json:"segment_ids"`
	TotalTimeSeconds int64           `db:"total_time_seconds" json:"total_time_seconds"`
	Efforts          json.RawMessage `db:"efforts" json:"efforts"`
}

type Map struct {
	ID              string    `db:"id" json:"id"`
	Polyline        string    `db:"polyline" json:"polyline"`
	SummaryPolyline string    `db:"summary_polyline" json:"summary_polyline"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type Segment struct {
	ID            int64   `db:"id" json:"id"`
	Name          string  `db:"name" json:"name"`
	ActivityType  string  `db:"activity_type" json:"activity_type"`
	Distance      float64 `db:"distance" json:"distance"`
	AverageGrade  float64 `db:"average_grade" json:"average_grade"`
	MaximumGrade  float64 `db:"maximum_grade" json:"maximum_grade"`
	ElevationHigh float64 `db:"elevation_high" json:"elevation_high"`
	ElevationLow  float64 `db:"elevation_low" json:"elevation_low"`
	StartLatlng   Floats  `db:"start_latlng" json:"start_latlng"`
	EndLatlng     Floats  `db:"end_latlng" json:"end_latlng"`
	// A small image of the elevation profile of this segment.
	ElevationProfile   string    `db:"elevation_profile" json:"elevation_profile"`
	ClimbCategory      int32     `db:"climb_category" json:"climb_category"`
	City               string    `db:"city" json:"city"`
	State              string    `db:"state" json:"state"`
	Country            string    `db:"country" json:"country"`
	Private            bool      `db:"private" json:"private"`
	Hazardous          bool      `db:"hazardous" json:"hazardous"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
	TotalElevationGain float64   `db:"total_elevation_gain" json:"total_elevation_gain"`
	MapID              string    `db:"map_id" json:"map_id"`
	TotalEffortCount   int32     `db:"total_effort_count" json:"total_effort_count"`
	TotalAthleteCount  int32     `db:"total_athlete_count" json:"total_athlete_count"`
	TotalStarCount     int32     `db:"total_star_count" json:"total_star_count"`
	// The time at which this segment was fetched from the Strava API.
	FetchedAt time.Time `db:"fetched_at" json:"fetched_at"`
	// Human friendly name for the segment
	FriendlyName string `db:"friendly_name" json:"friendly_name"`
}

type SegmentEffort struct {
	ID             int64     `db:"id" json:"id"`
	AthleteID      int64     `db:"athlete_id" json:"athlete_id"`
	SegmentID      int64     `db:"segment_id" json:"segment_id"`
	Name           string    `db:"name" json:"name"`
	ElapsedTime    float64   `db:"elapsed_time" json:"elapsed_time"`
	MovingTime     float64   `db:"moving_time" json:"moving_time"`
	StartDate      time.Time `db:"start_date" json:"start_date"`
	StartDateLocal time.Time `db:"start_date_local" json:"start_date_local"`
	// Distance is in meters
	Distance     float64       `db:"distance" json:"distance"`
	StartIndex   int32         `db:"start_index" json:"start_index"`
	EndIndex     int32         `db:"end_index" json:"end_index"`
	DeviceWatts  bool          `db:"device_watts" json:"device_watts"`
	AverageWatts float64       `db:"average_watts" json:"average_watts"`
	KomRank      sql.NullInt32 `db:"kom_rank" json:"kom_rank"`
	PrRank       sql.NullInt32 `db:"pr_rank" json:"pr_rank"`
	UpdatedAt    time.Time     `db:"updated_at" json:"updated_at"`
	// FK to activities table
	ActivitiesID int64 `db:"activities_id" json:"activities_id"`
}

type StarredSegment struct {
	AthleteID int64     `db:"athlete_id" json:"athlete_id"`
	SegmentID int64     `db:"segment_id" json:"segment_id"`
	Starred   bool      `db:"starred" json:"starred"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type SuperHugelActivity struct {
	AthleteID        int64           `db:"athlete_id" json:"athlete_id"`
	SegmentIds       interface{}     `db:"segment_ids" json:"segment_ids"`
	TotalTimeSeconds int64           `db:"total_time_seconds" json:"total_time_seconds"`
	Efforts          json.RawMessage `db:"efforts" json:"efforts"`
}

type WebhookDump struct {
	ID         uuid.UUID `db:"id" json:"id"`
	RecordedAt time.Time `db:"recorded_at" json:"recorded_at"`
	Raw        string    `db:"raw" json:"raw"`
}
