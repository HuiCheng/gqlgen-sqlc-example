-- name: ListDevices :many
SELECT
  device.*,
  sqlc.embed(device_auth)
FROM
  device
  JOIN device_auth ON device.auth_id = device_auth.id
WHERE
  (IIF(@ByAll, TRUE, NULL))
  OR (
    (
      IIF(@ByIDs, device.id IN (sqlc.slice ('IDs')), NULL)
    )
    OR (
      IIF(
        @ByAuthIDs,
        device.auth_id IN (sqlc.slice ('AuthIDs')),
        NULL
      )
    )
  )
LIMIT
  sqlc.arg (page_size)
OFFSET
  sqlc.arg (page_num);

-- name: CreateDevice :one 
INSERT INTO
  device (title, address, connect_type, auth_id, note)
VALUES
  (?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateDevice :exec
UPDATE device
SET
  title = ?,
  address = ?,
  connect_type = ?
WHERE
  id = ?;