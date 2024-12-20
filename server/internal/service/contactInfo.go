package service

import (
	"errors"
	"time"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type ContactInfoService struct {
	db *sqlx.DB
}

func NewContactInfoService(db *sqlx.DB) *ContactInfoService {
	return &ContactInfoService{
		db: db,
	}
}

func (s *ContactInfoService) GetByDriverId(driverId string, userId string) (*model.ContactInfo, error) {
	var contactInfo model.ContactInfo
	err := s.db.Get(&contactInfo,
		`SELECT c.* FROM contact_info c 
		JOIN drivers d ON c.id = d.contact_info_id 
		WHERE c.driver_id = $1 AND d.user_id = $2`,
		driverId,
		userId)
	if err != nil {
		return nil, err
	}

	return &contactInfo, nil
}

func (s *ContactInfoService) Create(driverId string, contactInfo model.ContactInfoInput) (*model.ContactInfo, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Queryx(
		`INSERT INTO contact_info (
            driver_id, phone_number, street_address, city, state, zip_code, country,
            created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING *`,
		driverId,
		contactInfo.PhoneNumber,
		contactInfo.StreetAddress,
		contactInfo.City,
		contactInfo.State,
		contactInfo.ZipCode,
		contactInfo.Country,
		time.Now(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	var result model.ContactInfo
	if !rows.Next() {
		return nil, errors.New("failed to create contact info")
	}
	if err := rows.StructScan(&result); err != nil {
		return nil, err
	}

	// Update driver with contact info ID
	_, err = tx.Exec(
		`UPDATE drivers 
		SET contact_info_id = $1
		WHERE id = $2`,
		result.Id,
		driverId,
	)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *ContactInfoService) Update(contactId string, userId string, input model.ContactInfoInput) (*model.ContactInfo, error) {
	rows, err := s.db.Queryx(
		`UPDATE contact_info c
		SET phone_number = $1,
			street_address = $2,
			city = $3,
			state = $4,
			zip_code = $5,
			country = $6,
			updated_at = NOW()
		FROM drivers d
		WHERE c.id = $7 
		AND c.driver_id = d.id 
		AND d.user_id = $8
		RETURNING c.*`,
		input.PhoneNumber,
		input.StreetAddress,
		input.City,
		input.State,
		input.ZipCode,
		input.Country,
		contactId,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var contactInfo model.ContactInfo
	if !rows.Next() {
		return nil, errors.New("contact info not found")
	}
	if err := rows.StructScan(&contactInfo); err != nil {
		return nil, err
	}

	return &contactInfo, nil
}

func (s *ContactInfoService) Delete(contactId string, userId string) (*model.ContactInfo, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// First update the driver to remove the reference
	_, err = tx.Exec(
		`UPDATE drivers d
		SET contact_info_id = NULL
		FROM contact_info c
		WHERE d.contact_info_id = c.id
		AND c.id = $1 
		AND d.user_id = $2`,
		contactId,
		userId,
	)
	if err != nil {
		return nil, err
	}

	// Then delete the contact info
	rows, err := tx.Queryx(
		`DELETE FROM contact_info c
		USING drivers d
		WHERE c.id = $1 
		AND c.driver_id = d.id 
		AND d.user_id = $2
		RETURNING c.*`,
		contactId,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var contactInfo model.ContactInfo
	if !rows.Next() {
		return nil, errors.New("contact info not found")
	}
	if err := rows.StructScan(&contactInfo); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &contactInfo, nil
}
