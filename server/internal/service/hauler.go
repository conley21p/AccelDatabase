package service

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type HaulerService struct {
	db *sqlx.DB
}

func NewHaulerService(db *sqlx.DB) *HaulerService {
	return &HaulerService{
		db: db,
	}
}

func (s *HaulerService) GetByDriverId(driverId string, userId string) (model.Hauler, error) {
	hauler := model.Hauler{}
	err := s.db.Get(&hauler,
		"select h.* from haulers h join drivers d on h.id = d.hauler_id where h.driver_id = $1 and d.user_id = $2",
		driverId,
		userId)
	if err != nil {
		return hauler, err
	}

	return hauler, nil
}

func (s *HaulerService) Create(driverId string, hauler model.HaulerInput) (*model.Hauler, error) {
	rows, err := s.db.Queryx(
		`INSERT INTO haulers (
            driver_id, make, model, year, mileage, towing_capacity
        ) VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, created_at, updated_at`,
		driverId,
		hauler.Make,
		hauler.Model,
		hauler.Year,
		hauler.Mileage,
		hauler.TowingCapacity,
	)
	if err != nil {
		return nil, err
	}

	var rtnhauler model.Hauler
	if !rows.Next() {
		return nil, errors.New("no hauler created")
	}
	if err := rows.StructScan(&rtnhauler); err != nil {
		return nil, err
	}

	return &rtnhauler, nil
}

func (s *HaulerService) Update(haulerId string, userId string, input model.HaulerInput) (*model.Hauler, error) {
	rows, err := s.db.Queryx(
		`UPDATE haulers h
		SET make = $1, 
			model = $2,
			year = $3,
			mileage = $4,
			towing_capacity = $5,
			updated_at = NOW()
		FROM drivers d
		WHERE h.id = $6 
		AND h.driver_id = d.id 
		AND d.user_id = $7
		RETURNING h.*`,
		input.Make,
		input.Model,
		input.Year,
		input.Mileage,
		input.TowingCapacity,
		haulerId,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var hauler model.Hauler
	if !rows.Next() {
		return nil, errors.New("hauler not found")
	}
	if err := rows.StructScan(&hauler); err != nil {
		return nil, err
	}

	return &hauler, nil
}

func (s *HaulerService) Delete(haulerId string, userId string) (*model.Hauler, error) {
	rows, err := s.db.Queryx(
		`DELETE FROM haulers h
		USING drivers d
		WHERE h.id = $1 
		AND h.driver_id = d.id 
		AND d.user_id = $2
		RETURNING h.*`,
		haulerId,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var hauler model.Hauler
	if !rows.Next() {
		return nil, errors.New("hauler not found")
	}
	if err := rows.StructScan(&hauler); err != nil {
		return nil, err
	}

	return &hauler, nil
}

func (s *HaulerService) GetAll(userId string) ([]model.Hauler, error) {
	var haulers []model.Hauler
	err := s.db.Select(&haulers,
		`SELECT h.* 
		FROM haulers h
		JOIN drivers d ON h.driver_id = d.id
		WHERE d.user_id = $1`,
		userId)
	if err != nil {
		return nil, err
	}

	return haulers, nil
}
