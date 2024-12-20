package service

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type BoatService struct {
	db *sqlx.DB
}

func NewBoatService(db *sqlx.DB) *BoatService {
	return &BoatService{
		db: db,
	}
}

func (s *BoatService) GetById(id string) (*model.Boat, error) {
	var boat model.Boat
	err := s.db.Get(&boat,
		`SELECT * FROM boats WHERE id = $1`,
		id)
	if err != nil {
		return nil, err
	}
	return &boat, nil
}

func (s *BoatService) Create(input model.BoatInput) (*model.Boat, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Queryx(
		`INSERT INTO boats (
			make,
			model,
			year,
			with_trailer,
			created_at
		) VALUES ($1, $2, $3, $4, NOW())
		RETURNING *`,
		input.Make,
		input.Model,
		input.Year,
		input.WithTrailer,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boat model.Boat
	if !rows.Next() {
		return nil, errors.New("failed to create boat")
	}
	if err := rows.StructScan(&boat); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &boat, nil
}

func (s *BoatService) Update(id string, input model.BoatInput) (*model.Boat, error) {
	rows, err := s.db.Queryx(
		`UPDATE boats 
		SET make = $1,
			model = $2,
			year = $3,
			with_trailer = $4,
			updated_at = NOW()
		WHERE id = $5
		RETURNING *`,
		input.Make,
		input.Model,
		input.Year,
		input.WithTrailer,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boat model.Boat
	if !rows.Next() {
		return nil, errors.New("boat not found")
	}
	if err := rows.StructScan(&boat); err != nil {
		return nil, err
	}

	return &boat, nil
}

func (s *BoatService) Delete(id string) (*model.Boat, error) {
	rows, err := s.db.Queryx(
		`DELETE FROM boats 
		WHERE id = $1
		RETURNING *`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var boat model.Boat
	if !rows.Next() {
		return nil, errors.New("boat not found")
	}
	if err := rows.StructScan(&boat); err != nil {
		return nil, err
	}

	return &boat, nil
}
