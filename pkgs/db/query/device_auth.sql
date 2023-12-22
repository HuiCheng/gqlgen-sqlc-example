-- name: GetDeviceAuth :one
SELECT
  *
FROM
  device_auth
WHERE
  id = ?;

-- name: ListDeviceAuth :many
SELECT
  device_auth.*,
  count(device.id) as device_count
FROM
  device_auth
  LEFT JOIN device ON device_auth.id = device.auth_id
WHERE
  (IIF(@ByAll, TRUE, NULL))
  OR (
    IIF(
      @ByIDs,
      device_auth.id IN (sqlc.slice ('ids')),
      NULL
    )
  )
GROUP BY
  device_auth.id
LIMIT
  @page_size
OFFSET
  @page_num;

-- name: CreateDeviceAuth :one
INSERT INTO
  device_auth (title, type, username, password, private_key)
VALUES
  (?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateDeviceAuth :one
UPDATE device_auth
SET
  title = ?,
  type = ?,
  username = ?,
  password = ?,
  private_key = ?,
  updated_at = DATETIME('now', 'localtime')
WHERE
  id = ? RETURNING *;

-- name: DeleteDeviceAuth :exec
DELETE FROM device_auth
Where
  id = ?;