-- name: GetUserData :one
SELECT * from user_data WHERE id = $1;

-- name: InsertUserData :one
INSERT INTO user_data (ip, user_agent, location, device, action, json_response_body, referrer, request_method, request_path) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;

-- name: GetAllUserData :many
SELECT * from user_data;
