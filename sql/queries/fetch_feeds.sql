-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at DESC LIMIT $1;

-- name: MarkFeedFetched :one
UPDATE feeds SET updated_at = $1, last_fetched_at = $1 WHERE id = $2
RETURNING *;