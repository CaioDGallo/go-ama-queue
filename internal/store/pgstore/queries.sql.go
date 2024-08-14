// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package pgstore

import (
	"context"

	"github.com/google/uuid"
)

const getAllUserData = `-- name: GetAllUserData :many
SELECT id, ip, user_agent, timestamp, location, device, action, json_response_body, referrer, request_method, request_path from user_data
`

func (q *Queries) GetAllUserData(ctx context.Context) ([]UserDatum, error) {
	rows, err := q.db.Query(ctx, getAllUserData)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserDatum
	for rows.Next() {
		var i UserDatum
		if err := rows.Scan(
			&i.ID,
			&i.Ip,
			&i.UserAgent,
			&i.Timestamp,
			&i.Location,
			&i.Device,
			&i.Action,
			&i.JsonResponseBody,
			&i.Referrer,
			&i.RequestMethod,
			&i.RequestPath,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserData = `-- name: GetUserData :one
SELECT id, ip, user_agent, timestamp, location, device, action, json_response_body, referrer, request_method, request_path from user_data WHERE id = $1
`

func (q *Queries) GetUserData(ctx context.Context, id uuid.UUID) (UserDatum, error) {
	row := q.db.QueryRow(ctx, getUserData, id)
	var i UserDatum
	err := row.Scan(
		&i.ID,
		&i.Ip,
		&i.UserAgent,
		&i.Timestamp,
		&i.Location,
		&i.Device,
		&i.Action,
		&i.JsonResponseBody,
		&i.Referrer,
		&i.RequestMethod,
		&i.RequestPath,
	)
	return i, err
}

const insertUserData = `-- name: InsertUserData :one
INSERT INTO user_data (ip, user_agent, location, device, action, json_response_body, referrer, request_method, request_path) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, ip, user_agent, timestamp, location, device, action, json_response_body, referrer, request_method, request_path
`

type InsertUserDataParams struct {
	Ip               string
	UserAgent        string
	Location         string
	Device           string
	Action           string
	JsonResponseBody string
	Referrer         string
	RequestMethod    string
	RequestPath      string
}

func (q *Queries) InsertUserData(ctx context.Context, arg InsertUserDataParams) (UserDatum, error) {
	row := q.db.QueryRow(ctx, insertUserData,
		arg.Ip,
		arg.UserAgent,
		arg.Location,
		arg.Device,
		arg.Action,
		arg.JsonResponseBody,
		arg.Referrer,
		arg.RequestMethod,
		arg.RequestPath,
	)
	var i UserDatum
	err := row.Scan(
		&i.ID,
		&i.Ip,
		&i.UserAgent,
		&i.Timestamp,
		&i.Location,
		&i.Device,
		&i.Action,
		&i.JsonResponseBody,
		&i.Referrer,
		&i.RequestMethod,
		&i.RequestPath,
	)
	return i, err
}