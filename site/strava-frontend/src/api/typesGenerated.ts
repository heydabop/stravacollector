
// Code generated by 'make site/strava-frontend/src/api/typesGenerated.ts'. DO NOT EDIT.

// From codersdk/athlete.go
export interface AthleteLogin {
  readonly athlete_id: string
  readonly summit: boolean
}

// From codersdk/athlete.go
export interface AthleteSummary {
  readonly athlete_id: string
  readonly summit: boolean
  readonly username: string
  readonly firstname: string
  readonly lastname: string
  readonly sex: string
  readonly profile_pic_link: string
  readonly profile_pic_link_medium: string
  readonly updated_at: string
  readonly hugel_count: number
}

// From codersdk/route.go
export interface CompetitiveRoute {
  readonly name: string
  readonly display_name: string
  readonly description: string
  readonly segments: SegmentSummary[]
}

// From codersdk/route.go
export interface DetailedSegment {
  readonly id: string
  readonly name: string
  readonly activity_type: string
  readonly distance: number
  readonly average_grade: number
  readonly maximum_grade: number
  readonly elevation_high: number
  readonly elevation_low: number
  readonly start_latlng: number[]
  readonly end_latlng: number[]
  readonly elevation_profile: string
  readonly climb_category: number
  readonly city: string
  readonly state: string
  readonly country: string
  readonly private: boolean
  readonly hazardous: boolean
  readonly created_at: string
  readonly updated_at: string
  readonly total_elevation_gain: number
  readonly map: Map
  readonly total_effort_count: number
  readonly total_athlete_count: number
  readonly total_star_count: number
  readonly fetched_at: string
}

// From codersdk/athlete.go
export interface HugelLeaderBoard {
  readonly personal_best?: HugelLeaderBoardActivity
  readonly activities: HugelLeaderBoardActivity[]
}

// From codersdk/athlete.go
export interface HugelLeaderBoardActivity {
  readonly rank_one_elapsed: number
  readonly activity_id: string
  readonly athlete_id: string
  readonly elapsed: number
  readonly rank: number
  readonly efforts: SegmentEffort[]
  readonly athlete: MinAthlete
  readonly activity_name: string
  readonly activity_distance: number
  readonly activity_moving_time: number
  readonly activity_elapsed_time: number
  readonly activity_start_date: string
  readonly activity_total_elevation_gain: number
}

// From codersdk/map.go
export interface Map {
  readonly id: string
  readonly polyline: string
  readonly summary_polyline: string
  readonly updated_at: string
}

// From codersdk/athlete.go
export interface MinAthlete {
  readonly athlete_id: string
  readonly username: string
  readonly firstname: string
  readonly lastname: string
  readonly sex: string
  readonly profile_pic_link: string
  readonly hugel_count: number
}

// From codersdk/response.go
export interface Response {
  readonly message: string
  readonly detail?: string
}

// From codersdk/athlete.go
export interface SegmentEffort {
  readonly activity_id: string
  readonly effort_id: string
  readonly start_date: string
  readonly segment_id: string
  readonly elapsed_time: number
  readonly moving_time: number
  readonly device_watts: boolean
  readonly average_watts: number
}

// From codersdk/route.go
export interface SegmentSummary {
  readonly id: string
  readonly name: string
}

// From codersdk/athlete.go
export interface SuperHugelLeaderBoard {
  readonly personal_best?: SuperHugelLeaderBoardActivity
  readonly activities: SuperHugelLeaderBoardActivity[]
}

// From codersdk/athlete.go
export interface SuperHugelLeaderBoardActivity {
  readonly rank_one_elapsed: number
  readonly athlete_id: string
  readonly elapsed: number
  readonly rank: number
  readonly efforts: SegmentEffort[]
  readonly athlete: MinAthlete
}

// From codersdk/int.go
export type StringInt = never
export const StringInts: StringInt[] = []
