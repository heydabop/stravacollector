// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package database

import "context"

const getAthletes = `-- name: GetAthletes :many
SELECT id, test FROM athletes
`

func (q *sqlQuerier) GetAthletes(ctx context.Context) ([]Athlete, error) {
	rows, err := q.db.QueryContext(ctx, getAthletes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Athlete
	for rows.Next() {
		var i Athlete
		if err := rows.Scan(&i.ID, &i.Test); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
