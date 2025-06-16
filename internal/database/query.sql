-- name: DeleteInquiry :execrows
DELETE FROM inquiries
WHERE id = ? ;

-- name: SelectInquiry :one
SELECT * FROM inquiries
WHERE id = ?
LIMIT 1 ;

-- name: SelectInquiries :many
SELECT * FROM inquiries
LIMIT ? ;
