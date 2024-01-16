// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tranfer.sql

package db

import (
	"context"
	"database/sql"
)

const createTransfers = `-- name: CreateTransfers :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount,
  created_at
) VALUES (
  $1, $2 , $3, $4
)
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransfersParams struct {
	FromAccountID sql.NullInt64 `json:"from_account_id"`
	ToAccountID   sql.NullInt64 `json:"to_account_id"`
	Amount        int64         `json:"amount"`
	CreatedAt     sql.NullTime  `json:"created_at"`
}

func (q *Queries) CreateTransfers(ctx context.Context, arg CreateTransfersParams) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, createTransfers,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
		arg.CreatedAt,
	)
	var i Transfers
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfers = `-- name: DeleteTransfers :exec
DELETE FROM transfers WHERE id = $1
`

func (q *Queries) DeleteTransfers(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfers, id)
	return err
}

const deleteTransfersByAccountId = `-- name: DeleteTransfersByAccountId :exec
DELETE FROM transfers WHERE from_account_id = $1 or to_account_id = $1
`

func (q *Queries) DeleteTransfersByAccountId(ctx context.Context, accountID sql.NullInt64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfersByAccountId, accountID)
	return err
}

const getTransfers = `-- name: GetTransfers :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfers(ctx context.Context, id int64) (Transfers, error) {
	row := q.db.QueryRowContext(ctx, getTransfers, id)
	var i Transfers
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
`

func (q *Queries) ListTransfers(ctx context.Context) ([]Transfers, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfers{}
	for rows.Next() {
		var i Transfers
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransfers = `-- name: UpdateTransfers :exec
UPDATE transfers SET amount = $2
WHERE id = $1
`

type UpdateTransfersParams struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

func (q *Queries) UpdateTransfers(ctx context.Context, arg UpdateTransfersParams) error {
	_, err := q.db.ExecContext(ctx, updateTransfers, arg.ID, arg.Amount)
	return err
}
