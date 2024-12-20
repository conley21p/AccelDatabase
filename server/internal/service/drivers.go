package service

import (
	"database/sql"
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type DriverService struct {
	db *sqlx.DB
}

func NewDriverService(db *sqlx.DB) *DriverService {
	return &DriverService{
		db: db,
	}
}
func (s *DriverService) GetByUserId(userId string) (*model.Driver, error) {
	driver := model.Driver{}
	err := s.db.Get(&driver, "select * from drivers where user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	driver, err = s.fetchDriverAttributes(driver)
	if err != nil {
		return nil, err
	}
	return &driver, nil
}

func (s *DriverService) GetById(id string) (*model.Driver, error) {
	driver := model.Driver{}
	println("Select on drivers")
	err := s.db.Get(&driver, "select * from drivers where id = $1", id)
	if err != nil {
		return nil, err
	}

	println("FetchDriverAttributes")
	driver, err = s.fetchDriverAttributes(driver)
	if err != nil {
		return nil, err
	}
	return &driver, nil
}

func (s *DriverService) fetchDriverAttributes(driver model.Driver) (model.Driver, error) {
	// Fetch ContactInfo
	contactInfo, err := s.fetchContactInfo(driver.Id)
	if err != nil && err != sql.ErrNoRows {
		return driver, err
	}
	driver.ContactInfo = contactInfo

	// Fetch Insurance
	insurance, err := s.fetchInsurance(driver.Id)
	if err != nil && err != sql.ErrNoRows {
		return driver, err
	}
	driver.Insurance = insurance

	// Fetch License
	license, err := s.fetchLicense(driver.Id)
	if err != nil && err != sql.ErrNoRows {
		return driver, err
	}
	driver.License = license

	// Fetch Rating
	rating, err := s.fetchRating(driver.Id)
	if err != nil && err != sql.ErrNoRows {
		return driver, err
	}
	driver.Rating = rating

	// Fetch Haulers
	haulers, err := s.fetchHaulersForDriver(driver.Id)
	if err != nil {
		return driver, err
	}
	driver.Haulers = haulers

	// Fetch Transportations
	transportations, err := s.fetchTransportationsForDriver(driver.Id)
	if err != nil {
		return driver, err
	}
	driver.Transportations = transportations
	// Fetch Offers
	offers, err := s.fetchOffersForDriver(driver.Id)
	if err != nil {
		return driver, err
	}
	driver.Offers = offers
	println("after select offers")
	return driver, nil
}

func (s *DriverService) fetchHaulersForDriver(driverId string) ([]model.Hauler, error) {
	var haulers []model.Hauler
	err := s.db.Select(&haulers,
		`SELECT * FROM haulers 
		WHERE driver_id = $1`,
		driverId)
	if err != nil {
		return haulers, err
	}

	// Fetch trailers for each hauler
	for i := range haulers {
		trailers, err := s.fetchTrailersForDriver(haulers[i].Id)
		if err != nil {
			return haulers, err
		}
		haulers[i].Trailers = trailers
	}
	return haulers, nil
}

func (s *DriverService) fetchTrailersForDriver(haulerId string) ([]model.Trailer, error) {
	var trailers []model.Trailer
	err := s.db.Select(&trailers, "select * from trailers where hauler_id = $1", haulerId)
	return trailers, err
}

func (s *DriverService) fetchTransportationsForDriver(driverId string) ([]model.Transportation, error) {
	var transportations []model.Transportation
	err := s.db.Select(&transportations, "select * from transportation where driver_id = $1", driverId)
	return transportations, err
}

func (s *DriverService) fetchOffersForDriver(driverId string) ([]model.Offer, error) {
	var offers []model.Offer
	err := s.db.Select(&offers, "select * from offers where driver_id = $1", driverId)
	return offers, err
}

func (s *DriverService) Create(userId string, driver model.DriverInput) (*model.Driver, error) {
	rows, err := s.db.Queryx(
		`INSERT INTO drivers (
            user_id, first_name, last_name, contact_info_id, insurance_id, license_id, rating_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id, created_at, updated_at`,
		userId,
		driver.FirstName,
		driver.LastName,
	)
	if err != nil {
		return nil, err
	}

	rtndriver := model.Driver{}

	rows.Next()
	err = rows.StructScan(&rtndriver)

	if err != nil {
		return nil, err
	}

	return &rtndriver, err
}

func (s *DriverService) Update(driverId string, userId string, input model.DriverInput) (*model.Driver, error) {
	rows, err := s.db.Queryx(
		`UPDATE drivers 
		SET first_name = $1, 
			last_name = $2,
			updated_at = NOW()
		WHERE id = $3 AND user_id = $4
		RETURNING *`,
		input.FirstName,
		input.LastName,
		driverId,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var driver model.Driver
	if rows.Next() {
		err = rows.StructScan(&driver)
		if err != nil {
			return nil, err
		}

		driver, err = s.fetchDriverAttributes(driver)
		if err != nil {
			return nil, err
		}
	}

	return &driver, nil
}

func (s *DriverService) Delete(driverId string, userId string) (*model.Driver, error) {
	rows, err := s.db.Queryx(
		`DELETE FROM drivers 
		WHERE id = $1 AND user_id = $2
		RETURNING *`,
		driverId,
		userId,
	)
	if err != nil {
		return nil, err
	}

	var driver model.Driver
	if !rows.Next() {
		return nil, errors.New("driver not found")
	}
	if err := rows.StructScan(&driver); err != nil {
		return nil, err
	}

	return &driver, nil
}

func (s *DriverService) createDriver(tx *sqlx.Tx, userId string, input model.DriverInput) (string, error) {
	driverRows, err := tx.Queryx(
		`INSERT INTO drivers (
			user_id, 
			first_name, 
			last_name, 
			created_at
		) VALUES (
			$1, 
			$2, 
			$3, 
			NOW()
		)
		RETURNING id`,
		userId,
		input.FirstName,
		input.LastName,
	)
	if err != nil {
		return "", err
	}
	defer driverRows.Close()

	var driverId string
	if !driverRows.Next() {
		return "", errors.New("failed to create driver")
	}
	if err := driverRows.Scan(&driverId); err != nil {
		return "", err
	}

	return driverId, nil
}

func (s *DriverService) createInsurance(tx *sqlx.Tx, driverId string, insurance *model.InsuranceInput) (*string, error) {
	if insurance == nil {
		return nil, nil
	}

	insuranceRows, err := tx.Queryx(
		`INSERT INTO insurance (
			driver_id, 
			policy_number, 
			ins_provider, 
			policy_start_date, 
			policy_end_date, 
			created_at
		) VALUES (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			NOW()
		)
		RETURNING id`,
		driverId,
		insurance.PolicyNumber,
		insurance.InsProvider,
		insurance.PolicyStartDate,
		insurance.PolicyEndDate,
	)
	if err != nil {
		return nil, err
	}
	defer insuranceRows.Close()

	var id string
	if insuranceRows.Next() {
		if err := insuranceRows.Scan(&id); err != nil {
			return nil, err
		}
		return &id, nil
	}
	return nil, nil
}

func (s *DriverService) createLicense(tx *sqlx.Tx, driverId string, license *model.LicenseInput) (*string, error) {
	if license == nil {
		return nil, nil
	}

	licRows, err := tx.Queryx(
		`INSERT INTO license (
			driver_id, 
			license_number, 
			license_expire_date, 
			created_at
		) VALUES ($1, $2, $3, NOW())
		RETURNING id`,
		driverId,
		license.LicenseNumber,
		license.LicenseExpireDate,
	)
	if err != nil {
		return nil, err
	}
	defer licRows.Close()

	var id string
	if licRows.Next() {
		if err := licRows.Scan(&id); err != nil {
			return nil, err
		}
		return &id, nil
	}
	return nil, nil
}

func (s *DriverService) updateDriverReferences(tx *sqlx.Tx, driverId string, insuranceId, licenseId *string) error {
	if insuranceId != nil || licenseId != nil {
		println("insuranceId: " + *insuranceId)
		println("licenseId: " + *licenseId)
		println("driverId: " + driverId)
		_, err := tx.Exec(
			`UPDATE drivers 
			SET insurance_id = $1, license_id = $2
			WHERE id = $3`,
			insuranceId,
			licenseId,
			driverId,
		)
		return err
	}
	return nil
}

func (s *DriverService) CreateWithDetails(userId string, input model.DriverRegistrationInput) (*model.Driver, error) {
	// Check if driver already exists for this user
	var count int
	err := s.db.Get(&count, "SELECT COUNT(*) FROM drivers WHERE user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("driver account already exists for this user")
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	println("CreateDriver")
	// Create driver
	driverId, err := s.createDriver(tx, userId, input.Driver)
	if err != nil {
		return nil, err
	}

	println("CreateInsurance")
	// Create insurance if provided
	insuranceId, err := s.createInsurance(tx, driverId, input.Insurance)
	if err != nil {
		return nil, err
	}

	println("CreateLicense")
	// Create license if provided
	licenseId, err := s.createLicense(tx, driverId, input.License)
	if err != nil {
		return nil, err
	}

	println("UpdateDriverReferences")
	// Update driver with references
	if err := s.updateDriverReferences(tx, driverId, insuranceId, licenseId); err != nil {
		return nil, err
	}

	println("Commit")
	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	println("GetById")
	// Fetch the complete driver record
	return s.GetById(driverId)
}

func (s *DriverService) fetchContactInfo(driverId string) (*model.ContactInfo, error) {
	var contactInfo model.ContactInfo
	println("driverId: " + driverId)
	err := s.db.Get(&contactInfo,
		`SELECT c.* FROM contact_info c 
		JOIN drivers d ON c.id = d.contact_info_id 
		WHERE d.id = $1`,
		driverId)
	if err != nil {
		return nil, err
	}
	println("after select contactInfo")
	return &contactInfo, nil
}

func (s *DriverService) fetchInsurance(driverId string) (*model.Insurance, error) {
	var insurance model.Insurance
	err := s.db.Get(&insurance,
		`SELECT i.* FROM insurance i 
		JOIN drivers d ON i.id = d.insurance_id 
		WHERE d.id = $1`,
		driverId)
	if err != nil {
		return nil, err
	}
	return &insurance, nil
}

func (s *DriverService) fetchLicense(driverId string) (*model.License, error) {
	var license model.License
	err := s.db.Get(&license,
		`SELECT l.* FROM license l 
		JOIN drivers d ON l.id = d.license_id 
		WHERE d.id = $1`,
		driverId)
	if err != nil {
		return nil, err
	}
	return &license, nil
}

func (s *DriverService) fetchRating(driverId string) (*model.Rating, error) {
	var rating model.Rating
	err := s.db.Get(&rating,
		`SELECT r.* FROM ratings r 
		JOIN drivers d ON r.id = d.rating_id 
		WHERE d.id = $1`,
		driverId)
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

func (s *DriverService) GetWithDetails(driverId string) (*model.Driver, error) {
	driver := model.Driver{}
	err := s.db.Get(&driver, "SELECT * FROM drivers WHERE id = $1", driverId)
	if err != nil {
		return nil, err
	}

	driver, err = s.fetchDriverAttributes(driver)
	if err != nil {
		return nil, err
	}

	return &driver, nil
}
