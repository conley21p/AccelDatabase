package service

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type OwnerService struct {
	db *sqlx.DB
}

func NewOwnerService(db *sqlx.DB) *OwnerService {
	return &OwnerService{
		db: db,
	}
}

func (s *OwnerService) GetByUserId(userId string) (*model.Owner, error) {
	var owner model.Owner
	err := s.db.Get(&owner,
		`SELECT * FROM owners WHERE user_id = $1`,
		userId)
	if err != nil {
		return nil, err
	}
	return &owner, nil
}

func (s *OwnerService) Create(userId string, input model.OwnerInput) (*model.Owner, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Queryx(
		`INSERT INTO owners (
			user_id,
			first_name,
			last_name,
			created_at
		) VALUES ($1, $2, $3, NOW())
		RETURNING *`,
		userId,
		input.FirstName,
		input.LastName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var owner model.Owner
	if !rows.Next() {
		return nil, errors.New("failed to create owner")
	}
	if err := rows.StructScan(&owner); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &owner, nil
}

func (s *OwnerService) Update(ownerId string, userId string, input model.OwnerInput) (*model.Owner, error) {
	rows, err := s.db.Queryx(
		`UPDATE owners 
		SET first_name = $1,
			last_name = $2,
			updated_at = NOW()
		WHERE id = $3 AND user_id = $4
		RETURNING *`,
		input.FirstName,
		input.LastName,
		ownerId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var owner model.Owner
	if !rows.Next() {
		return nil, errors.New("owner not found")
	}
	if err := rows.StructScan(&owner); err != nil {
		return nil, err
	}

	return &owner, nil
}

func (s *OwnerService) Delete(ownerId string, userId string) (*model.Owner, error) {
	rows, err := s.db.Queryx(
		`DELETE FROM owners 
		WHERE id = $1 AND user_id = $2
		RETURNING *`,
		ownerId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var owner model.Owner
	if !rows.Next() {
		return nil, errors.New("owner not found")
	}
	if err := rows.StructScan(&owner); err != nil {
		return nil, err
	}

	return &owner, nil
}
