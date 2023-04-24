-- name: DeleteActivity :one
DELETE FROM
	activities
WHERE
	id = $1
RETURNING *
;

-- name: GetActivity :one
SELECT
	*
FROM
    activities
WHERE
    id = $1;

-- name: UpdateActivityName :exec
UPDATE activities
SET
    name = $2
WHERE
    id = $1;

-- name: UpsertActivity :one
INSERT INTO
	activities(
	updated_at,
	id, athlete_id, upload_id, external_id, name,
	moving_time, elapsed_time, total_elevation_gain, activity_type,
	sport_type, start_date, start_date_local, timezone, utc_offset,
	start_latlng, end_latlng, achievement_count, kudos_count,
	comment_count, athlete_count, photo_count, map_id, map_polyline,
	map_summary_polyline, trainer, commute, manual, private, flagged,
	gear_id, from_accepted_tag, average_speed, max_speed, average_cadence,
	average_temp, average_watts, weighted_average_watts, kilojoules,
	device_watts, has_heartrate, max_watts, elev_high, elev_low, pr_count,
	total_photo_count, workout_type, suffer_score, calories,
	embed_token, segment_leaderboard_opt_out, leaderboard_opt_out,
	num_segment_efforts, premium_fetch

)
VALUES
	(Now(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,
	 $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33,
	 $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49,
	 $50, $51, $52, $53)
ON CONFLICT
	(id)
	DO UPDATE SET
	updated_at = Now(),
	athlete_id = $2,
	upload_id = $3,
	external_id = $4,
	name = $5,
	moving_time = $6,
	elapsed_time = $7,
	total_elevation_gain = $8,
	activity_type = $9,
	sport_type = $10,
	start_date = $11,
	start_date_local = $12,
	timezone = $13,
	utc_offset = $14,
	start_latlng = $15,
	end_latlng = $16,
	achievement_count = $17,
	kudos_count = $18,
	comment_count = $19,
	athlete_count = $20,
	photo_count = $21,
	map_id = $22,
	map_polyline = $23,
	map_summary_polyline = $24,
	trainer = $25,
	commute = $26,
	manual = $27,
	private = $28,
	flagged = $29,
	gear_id = $30,
	from_accepted_tag = $31,
	average_speed = $32,
	max_speed = $33,
	average_cadence = $34,
	average_temp = $35,
	average_watts = $36,
	weighted_average_watts = $37,
	kilojoules = $38,
	device_watts = $39,
	has_heartrate = $40,
	max_watts = $41,
	elev_high = $42,
	elev_low = $43,
	pr_count = $44,
	total_photo_count = $45,
	workout_type = $46,
	suffer_score = $47,
	calories = $48,
	embed_token = $49,
	segment_leaderboard_opt_out = $50,
	leaderboard_opt_out = $51,
	num_segment_efforts = $52,
	premium_fetch = $53
RETURNING *;