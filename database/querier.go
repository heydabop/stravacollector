// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package database

import (
	"context"
)

type sqlcQuerier interface {
	GetAthletes(ctx context.Context) ([]Athlete, error)
}

var _ sqlcQuerier = (*sqlQuerier)(nil)
