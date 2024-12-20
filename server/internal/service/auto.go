package service

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type AutoService struct {
	db *sqlx.DB
}

func NewAutoService(db *sqlx.DB) *AutoService {
	return &AutoService{
		db: db,
	}
}

func (s *AutoService) GetById(id string) (*model.Auto, error) {
	var auto model.Auto
	err := s.db.Get(&auto,
		`SELECT * FROM autos WHERE id = $1`,
		id)
	if err != nil {
		return nil, err
	}
	return &auto, nil
}

func (s *AutoService) Create(input model.AutoInput) (*model.Auto, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Queryx(
		`INSERT INTO autos (
			make,
			model,
			year,
			created_at
		) VALUES ($1, $2, $3, NOW())
		RETURNING *`,
		input.Make,
		input.Model,
		input.Year,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auto model.Auto
	if !rows.Next() {
		return nil, errors.New("failed to create auto")
	}
	if err := rows.StructScan(&auto); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &auto, nil
}

func (s *AutoService) Update(id string, input model.AutoInput) (*model.Auto, error) {
	rows, err := s.db.Queryx(
		`UPDATE autos 
		SET make = $1,
			model = $2,
			year = $3,
			updated_at = NOW()
		WHERE id = $4
		RETURNING *`,
		input.Make,
		input.Model,
		input.Year,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auto model.Auto
	if !rows.Next() {
		return nil, errors.New("auto not found")
	}
	if err := rows.StructScan(&auto); err != nil {
		return nil, err
	}

	return &auto, nil
}

func (s *AutoService) Delete(id string) (*model.Auto, error) {
	rows, err := s.db.Queryx(
		`DELETE FROM autos 
		WHERE id = $1
		RETURNING *`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auto model.Auto
	if !rows.Next() {
		return nil, errors.New("auto not found")
	}
	if err := rows.StructScan(&auto); err != nil {
		return nil, err
	}

	return &auto, nil
}
