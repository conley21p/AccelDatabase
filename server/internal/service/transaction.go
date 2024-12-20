package service

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type TransactionService struct {
	db *sqlx.DB
}

func NewTransactionService(db *sqlx.DB) *TransactionService {
	return &TransactionService{
		db: db,
	}
}

func (s *TransactionService) GetById(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := s.db.Get(&transaction,
		`SELECT * FROM transactions WHERE id = $1`,
		id)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (s *TransactionService) Create(input model.TransactionInput) (*model.Transaction, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Queryx(
		`INSERT INTO transactions (
			transportation_id,
			driver_id,
			buyer_id,
			payment_method,
			amount,
			transaction_date,
			created_at
		) VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING *`,
		input.TransportationId,
		input.DriverId,
		input.BuyerId,
		input.PaymentMethod,
		input.Amount,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transaction model.Transaction
	if !rows.Next() {
		return nil, errors.New("failed to create transaction")
	}
	if err := rows.StructScan(&transaction); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (s *TransactionService) Update(id string, input model.TransactionInput) (*model.Transaction, error) {
	rows, err := s.db.Queryx(
		`UPDATE transactions 
		SET transportation_id = $1,
			driver_id = $2,
			buyer_id = $3,
			payment_method = $4,
			amount = $5,
			updated_at = NOW()
		WHERE id = $6
		RETURNING *`,
		input.TransportationId,
		input.DriverId,
		input.BuyerId,
		input.PaymentMethod,
		input.Amount,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transaction model.Transaction
	if !rows.Next() {
		return nil, errors.New("transaction not found")
	}
	if err := rows.StructScan(&transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (s *TransactionService) Delete(id string) (*model.Transaction, error) {
	rows, err := s.db.Queryx(
		`DELETE FROM transactions 
		WHERE id = $1
		RETURNING *`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transaction model.Transaction
	if !rows.Next() {
		return nil, errors.New("transaction not found")
	}
	if err := rows.StructScan(&transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}
