-- name: InsertBooks :exec
INSERT INTO books (
    name,
    release_year,
    total_page
) VALUES (
    ?, ?, ?
)
RETURNING *;

-- name: FullScanOfBooks :many
SELECT *
FROM books;