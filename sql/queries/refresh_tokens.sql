-- name: CreateRefreshToken :one
INSERT INTO
  refresh_tokens (
    token,
    created_at,
    updated_at,
    user_id,
    expires_at
  )
VALUES
  ($1, NOW(), NOW(), $2, $3)
RETURNING
  *;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET
  updated_at = NOW(),
  revoked_at = NOW()
WHERE
  token = $1;
