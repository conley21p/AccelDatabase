package service

import (
	"errors"
	"time"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type TransportationService struct {
	db *sqlx.DB
}

func NewTransportationService(db *sqlx.DB) *TransportationService {
	return &TransportationService{
		db: db,
	}
}

func (s *TransportationService) GetById(id string) (*model.Transportation, error) {
	var transportation model.Transportation
	err := s.db.Get(&transportation,
		`SELECT * FROM transportation WHERE id = $1`,
		id)
	if err != nil {
		return nil, err
	}
	return &transportation, nil
}

func (s *TransportationService) Create(input model.TransportationInput) (*model.Transportation, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	const dateFormat = "2006-01-02"
	transportDate, err := time.Parse(dateFormat, input.TransportDate)
	if err != nil {
		return nil, errors.New("invalid transport date format")
	}

	deliverByDate, err := time.Parse(dateFormat, input.DeliverByDate)
	if err != nil {
		return nil, errors.New("invalid deliver by date format")
	}

	pickupByDate, err := time.Parse(dateFormat, input.PickupByDate)
	if err != nil {
		return nil, errors.New("invalid pickup by date format")
	}

	pickupAvailableDate, err := time.Parse(dateFormat, input.PickupAvailableDate)
	if err != nil {
		return nil, errors.New("invalid pickup available date format")
	}

	rows, err := tx.Queryx(
		`INSERT INTO transportation (
			description,
			transport_date,
			pickup_address,
			delivery_address,
			deliver_by_date,
			pickup_by_date,
			pickup_available_date,
			request_price,
			vehicle_id,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
		RETURNING *`,
		input.Description,
		transportDate,
		input.PickupAddress,
		input.DeliveryAddress,
		deliverByDate,
		pickupByDate,
		pickupAvailableDate,
		input.RequestPrice,
		input.VehicleId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transportation model.Transportation
	if !rows.Next() {
		return nil, errors.New("failed to create transportation")
	}
	if err := rows.StructScan(&transportation); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &transportation, nil
}

func (s *TransportationService) Update(id string, input model.TransportationInput) (*model.Transportation, error) {
	transportDate, err := time.Parse("2006-01-02", input.TransportDate)
	if err != nil {
		return nil, errors.New("invalid transport date format")
	}

	deliverByDate, err := time.Parse("2006-01-02", input.DeliverByDate)
	if err != nil {
		return nil, errors.New("invalid deliver by date format")
	}

	pickupByDate, err := time.Parse("2006-01-02", input.PickupByDate)
	if err != nil {
		return nil, errors.New("invalid pickup by date format")
	}

	pickupAvailableDate, err := time.Parse("2006-01-02", input.PickupAvailableDate)
	if err != nil {
		return nil, errors.New("invalid pickup available date format")
	}

	rows, err := s.db.Queryx(
		`UPDATE transportation 
		SET description = $1,
			transport_date = $2,
			pickup_address = $3,
			delivery_address = $4,
			deliver_by_date = $5,
			pickup_by_date = $6,
			pickup_available_date = $7,
			request_price = $8,
			vehicle_id = $9,
			updated_at = NOW()
		WHERE id = $10
		RETURNING *`,
		input.Description,
		transportDate,
		input.PickupAddress,
		input.DeliveryAddress,
		deliverByDate,
		pickupByDate,
		pickupAvailableDate,
		input.RequestPrice,
		input.VehicleId,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transportation model.Transportation
	if !rows.Next() {
		return nil, errors.New("transportation not found")
	}
	if err := rows.StructScan(&transportation); err != nil {
		return nil, err
	}

	return &transportation, nil
}

func (s *TransportationService) Delete(id string) (*model.Transportation, error) {
	rows, err := s.db.Queryx(
		`DELETE FROM transportation 
		WHERE id = $1
		RETURNING *`,
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transportation model.Transportation
	if !rows.Next() {
		return nil, errors.New("transportation not found")
	}
	if err := rows.StructScan(&transportation); err != nil {
		return nil, err
	}

	return &transportation, nil
}
