package service

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type LicenseService struct {
	db *sqlx.DB
}

func NewLicenseService(db *sqlx.DB) *LicenseService {
	return &LicenseService{
		db: db,
	}
}

func (s *LicenseService) GetByDriverId(driverId string, userId string) (*model.License, error) {
	var license model.License
	err := s.db.Get(&license,
		`SELECT l.* FROM license l 
		JOIN drivers d ON l.id = d.license_id 
		WHERE l.driver_id = $1 AND d.user_id = $2`,
		driverId,
		userId)
	if err != nil {
		return nil, err
	}
	return &license, nil
}

func (s *LicenseService) Create(driverId string, license model.LicenseInput) (*model.License, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Queryx(
		`INSERT INTO license (
			driver_id,
			license_number,
			license_expire_date,
			created_at
		) VALUES ($1, $2, $3, NOW())
		RETURNING *`,
		driverId,
		license.LicenseNumber,
		license.LicenseExpireDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.License
	if !rows.Next() {
		return nil, errors.New("failed to create license")
	}
	if err := rows.StructScan(&result); err != nil {
		return nil, err
	}

	// Update driver with license ID
	_, err = tx.Exec(
		`UPDATE drivers 
		SET license_id = $1
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

func (s *LicenseService) Update(licenseId string, userId string, input model.LicenseInput) (*model.License, error) {
	rows, err := s.db.Queryx(
		`UPDATE license l
		SET license_number = $1,
			license_expire_date = $2,
			updated_at = NOW()
		FROM drivers d
		WHERE l.id = $3 
		AND l.driver_id = d.id 
		AND d.user_id = $4
		RETURNING l.*`,
		input.LicenseNumber,
		input.LicenseExpireDate,
		licenseId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var license model.License
	if !rows.Next() {
		return nil, errors.New("license not found")
	}
	if err := rows.StructScan(&license); err != nil {
		return nil, err
	}

	return &license, nil
}

func (s *LicenseService) Delete(licenseId string, userId string) (*model.License, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// First update the driver to remove the reference
	_, err = tx.Exec(
		`UPDATE drivers d
		SET license_id = NULL
		FROM license l
		WHERE d.license_id = l.id
		AND l.id = $1 
		AND d.user_id = $2`,
		licenseId,
		userId,
	)
	if err != nil {
		return nil, err
	}

	// Then delete the license
	rows, err := tx.Queryx(
		`DELETE FROM license l
		USING drivers d
		WHERE l.id = $1 
		AND l.driver_id = d.id 
		AND d.user_id = $2
		RETURNING l.*`,
		licenseId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var license model.License
	if !rows.Next() {
		return nil, errors.New("license not found")
	}
	if err := rows.StructScan(&license); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &license, nil
}
