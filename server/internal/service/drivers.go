package service

import (
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
func (s *DriverService) GetByUserId(userId string) (model.Driver, error) {
	driver := model.Driver{}
	err := s.db.Get(&driver, "select * from drivers where user_id = $1", userId)
	if err != nil {
		return driver, err
	}

	driver, err = s.fetchDriverAttributes(driver)

	return driver, nil
}

func (s *DriverService) GetById(id string) (model.Driver, error) {
	driver := model.Driver{}
	err := s.db.Get(&driver, "select * from drivers where id = $1", id)
	if err != nil {
		return driver, err
	}

	driver, err = s.fetchDriverAttributes(driver)

	return driver, nil
}

func (s *DriverService) fetchDriverAttributes(driver model.Driver) (model.Driver, error) {
	// Fetch Haulers in a separate query
	haulers, err := s.fetchHaulersForDriver(driver.Id)
	if err != nil {
		return driver, err
	}
	driver.Haulers = haulers

	// Fetch Transportations in a separate query
	transportations, err := s.fetchTransportationsForDriver(driver.Id)
	if err != nil {
		return driver, err
	}
	driver.Transportations = transportations

	// Fetch Haulers in a separate query
	offers, err := s.fetchOffersForDriver(driver.Id)
	if err != nil {
		return driver, err
	}
	driver.Offers = offers

	return driver, err
}

func (s *DriverService) fetchHaulersForDriver(driverId string) ([]model.Hauler, error) {
	var haulers []model.Hauler
	err := s.db.Select(&haulers, "select * from haulers where driver_id = $1", driverId)
	if err != nil {
		return haulers, err
	}
	for _, hauler := range haulers {

		// Fetch Haulers in a separate query
		trailers, err := s.fetchTrailersForDriver(hauler.Id)
		if err != nil {
			return haulers, err
		}
		hauler.Trailers = trailers
	}

	return haulers, err
}

func (s *DriverService) fetchTrailersForDriver(haulerId string) ([]model.Trailer, error) {
	var trailers []model.Trailer
	err := s.db.Select(&trailers, "select * from trailers where hauler_id = $1", haulerId)
	return trailers, err
}

func (s *DriverService) fetchTransportationsForDriver(driverId string) ([]model.Transportation, error) {
	var transportations []model.Transportation
	err := s.db.Select(&transportations, "select * from transportations where driver_id = $1", driverId)
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
