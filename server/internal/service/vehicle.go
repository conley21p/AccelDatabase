package service

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type VehicleService struct {
	db *sqlx.DB
}

func NewVehicleService(db *sqlx.DB) *VehicleService {
	return &VehicleService{
		db: db,
	}
}

func (s *VehicleService) GetById(id string) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	err := s.db.Get(&vehicle,
		`SELECT * FROM vehicles WHERE id = $1`,
		id)
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (s *VehicleService) Create(input model.VehicleInput) (*model.Vehicle, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Queryx(
		`INSERT INTO vehicles (
			length,
			width,
			height,
			transportation_id,
			auto_id,
			boat_id,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING *`,
		input.Length,
		input.Width,
		input.Height,
		input.TransportationId,
		input.AutoId,
		input.BoatId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicle model.Vehicle
	if !rows.Next() {
		return nil, errors.New("failed to create vehicle")
	}
	if err := rows.StructScan(&vehicle); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (s *VehicleService) Update(id string, input model.VehicleInput) (*model.Vehicle, error) {
	rows, err := s.db.Queryx(
		`UPDATE vehicles 
		SET length = $1,
			width = $2,
			height = $3,
			transportation_id = $4,
			auto_id = $5,
			boat_id = $6,
			updated_at = NOW()
		WHERE id = $7
		RETURNING *`,
		input.Length,
		input.Width,
		input.Height,
		input.TransportationId,
		input.AutoId,
		input.BoatId,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicle model.Vehicle
	if !rows.Next() {
		return nil, errors.New("vehicle not found")
	}
	if err := rows.StructScan(&vehicle); err != nil {
		return nil, err
	}

	return &vehicle, nil
}

func (s *VehicleService) Delete(id string) (*model.Vehicle, error) {
	rows, err := s.db.Queryx(
		`DELETE FROM vehicles 
		WHERE id = $1
		RETURNING *`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicle model.Vehicle
	if !rows.Next() {
		return nil, errors.New("vehicle not found")
	}
	if err := rows.StructScan(&vehicle); err != nil {
		return nil, err
	}

	return &vehicle, nil
}
