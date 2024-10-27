// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: follows.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFollow = `-- name: CreateFollow :exec
INSERT INTO follows (
    following_user_id,
    followed_user_id
) VALUES (
    $1, $2
)
`

type CreateFollowParams struct {
	FollowingUserID pgtype.Int4
	FollowedUserID  pgtype.Int4
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) error {
	_, err := q.db.Exec(ctx, createFollow, arg.FollowingUserID, arg.FollowedUserID)
	return err
}

const deleteFollow = `-- name: DeleteFollow :exec
DELETE FROM follows
WHERE followed_user_id = $1
`

func (q *Queries) DeleteFollow(ctx context.Context, followedUserID pgtype.Int4) error {
	_, err := q.db.Exec(ctx, deleteFollow, followedUserID)
	return err
}

const getFollow = `-- name: GetFollow :one
SELECT following_user_id, followed_user_id, created_at FROM follows
RETURING
`

func (q *Queries) GetFollow(ctx context.Context) (Follow, error) {
	row := q.db.QueryRow(ctx, getFollow)
	var i Follow
	err := row.Scan(&i.FollowingUserID, &i.FollowedUserID, &i.CreatedAt)
	return i, err
}