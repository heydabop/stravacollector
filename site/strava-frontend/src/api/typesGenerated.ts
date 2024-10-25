// Code generated by 'make site/strava-frontend/src/api/typesGenerated.ts'. DO NOT EDIT.

// The code below is generated from api/modelsdk.

// From modelsdk/athlete.go
export interface ActivitySummary {
	readonly activity_id: string;
	readonly athlete_id: string;
	readonly upload_id: string;
	readonly external_id: string;
	readonly name: string;
	readonly distance: number;
	readonly moving_time: number;
	readonly elapsed_time: number;
	readonly total_elevation_gain: number;
	readonly activity_type: string;
	readonly sport_type: string;
	readonly start_date: string;
	readonly start_date_local: string;
	readonly timezone: string;
}

// From modelsdk/athlete.go
export interface AthleteHugelActivities {
	readonly activities: Readonly<Array<AthleteHugelActivity>>;
}

// From modelsdk/athlete.go
export interface AthleteHugelActivity {
	readonly summary: ActivitySummary;
	readonly efforts: Readonly<Array<SegmentEffort>>;
	readonly total_time_seconds: number;
}

// From modelsdk/athlete.go
export interface AthleteLoad {
	readonly athlete_id: number;
	readonly last_backload_activity_start: string;
	readonly last_load_attempt: string;
	readonly last_load_incomplete: boolean;
	readonly last_load_error: string;
	readonly activites_loaded_last_attempt: number;
	readonly earliest_activity: string;
	readonly earliest_activity_id: number;
	readonly earliest_activity_done: boolean;
}

// From modelsdk/athlete.go
export interface AthleteLogin {
	readonly athlete_id: string;
	readonly summit: boolean;
}

// From modelsdk/athlete.go
export interface AthleteSummary {
	readonly athlete_id: string;
	readonly summit: boolean;
	readonly username: string;
	readonly firstname: string;
	readonly lastname: string;
	readonly sex: string;
	readonly profile_pic_link: string;
	readonly profile_pic_link_medium: string;
	readonly updated_at: string;
	readonly hugel_count: number;
}

// From modelsdk/athlete.go
export interface AthleteSyncSummary {
	readonly athlete_load: AthleteLoad;
	readonly total_activities: number;
	readonly synced_activities: Readonly<Array<SyncActivitySummary>>;
	readonly athlete_summary: AthleteSummary;
	readonly total_summary: number;
	readonly total_detail: number;
}

// From modelsdk/route.go
export interface CompetitiveRoute {
	readonly name: string;
	readonly display_name: string;
	readonly description: string;
	readonly segments: Readonly<Array<SegmentSummary>>;
}

// From modelsdk/route.go
export interface DetailedSegment {
	readonly id: string;
	readonly name: string;
	readonly activity_type: string;
	readonly distance: number;
	readonly average_grade: number;
	readonly maximum_grade: number;
	readonly elevation_high: number;
	readonly elevation_low: number;
	readonly start_latlng: Readonly<Array<number>>;
	readonly end_latlng: Readonly<Array<number>>;
	readonly elevation_profile: string;
	readonly climb_category: number;
	readonly city: string;
	readonly state: string;
	readonly country: string;
	readonly private: boolean;
	readonly hazardous: boolean;
	readonly created_at: string;
	readonly updated_at: string;
	readonly total_elevation_gain: number;
	readonly map: Map;
	readonly total_effort_count: number;
	readonly total_athlete_count: number;
	readonly total_star_count: number;
	readonly fetched_at: string;
}

// From modelsdk/athlete.go
export interface HugelLeaderBoard {
	readonly personal_best?: HugelLeaderBoardActivity;
	readonly superlatives: List;
	readonly activities: Readonly<Array<HugelLeaderBoardActivity>>;
}

// From modelsdk/athlete.go
export interface HugelLeaderBoardActivity {
	readonly rank_one_elapsed: number;
	readonly activity_id: string;
	readonly athlete_id: string;
	readonly elapsed: number;
	readonly rank: number;
	readonly efforts: Readonly<Array<SegmentEffort>>;
	readonly athlete: MinAthlete;
	readonly activity_name: string;
	readonly activity_distance: number;
	readonly activity_moving_time: number;
	readonly activity_elapsed_time: number;
	readonly activity_start_date: string;
	readonly activity_total_elevation_gain: number;
	readonly activity_suffer_score: number;
	readonly activity_achievement_count: number;
}

// From modelsdk/map.go
export interface Map {
	readonly id: string;
	readonly polyline: string;
	readonly summary_polyline: string;
	readonly updated_at: string;
}

// From modelsdk/athlete.go
export interface MinAthlete {
	readonly athlete_id: string;
	readonly username: string;
	readonly firstname: string;
	readonly lastname: string;
	readonly sex: string;
	readonly profile_pic_link: string;
	readonly hugel_count: number;
}

// From modelsdk/route.go
export interface PersonalBestSegmentEffort {
	readonly best_effort_id: string;
	readonly best_effort_elapsed_time: number;
	readonly best_effort_moving_time: number;
	readonly best_effort_start_date: string;
	readonly best_effort_start_date_local: string;
	readonly best_effort_device_watts: boolean;
	readonly best_effort_average_watts: number;
	readonly best_effort_activities_id: string;
}

// From modelsdk/route.go
export interface PersonalSegment {
	readonly detailed_segment: DetailedSegment;
	readonly starred?: boolean;
	readonly personal_best?: PersonalBestSegmentEffort;
}

// From modelsdk/response.go
export interface Response {
	readonly message: string;
	readonly detail?: string;
}

// From modelsdk/athlete.go
export interface SegmentEffort {
	readonly activity_id: string;
	readonly effort_id: string;
	readonly start_date: string;
	readonly segment_id: string;
	readonly elapsed_time: number;
	readonly moving_time: number;
	readonly device_watts: boolean;
	readonly average_watts: number;
}

// From modelsdk/route.go
export interface SegmentSummary {
	readonly id: string;
	readonly name: string;
}

// From modelsdk/athlete.go
export interface SuperHugelLeaderBoard {
	readonly personal_best?: SuperHugelLeaderBoardActivity;
	readonly activities: Readonly<Array<SuperHugelLeaderBoardActivity>>;
}

// From modelsdk/athlete.go
export interface SuperHugelLeaderBoardActivity {
	readonly rank_one_elapsed: number;
	readonly athlete_id: string;
	readonly elapsed: number;
	readonly rank: number;
	readonly efforts: Readonly<Array<SegmentEffort>>;
	readonly athlete: MinAthlete;
}

// From modelsdk/athlete.go
export interface SyncActivitySummary {
	readonly activity_summary: ActivitySummary;
	readonly synced: boolean;
	readonly synced_at: string;
}

// From modelsdk/route.go
export interface VerifyRouteResponse {
	readonly missing_segments: Readonly<Array<SegmentSummary>>;
}

// From modelsdk/int.go
export type StringInt = never
export const StringInts: StringInt[] = []

// The code below is generated from api/modelsdk/sdktype.



// The code below is generated from api/superlative.

// From superlative/superlative.go
export interface Entry<T extends comparable> {
	readonly activity_id: string;
	readonly value: T;
}

// From superlative/superlative.go
export interface List {
	readonly early_bird: Entry<string>;
	readonly night_owl: Entry<string>;
	readonly most_stoppage: Entry<number>;
	readonly least_stoppage: Entry<number>;
	readonly most_watts: Entry<number>;
	readonly most_cadence: Entry<number>;
	readonly least_cadence: Entry<number>;
	readonly most_suffer: Entry<number>;
	readonly most_achievements: Entry<number>;
	readonly longest_ride: Entry<number>;
	readonly shortest_ride: Entry<number>;
}

export type comparable = boolean | number | string | any

