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

func (s *DriverService) GetById(id string) (model.Driver, error) {
	driver := model.Driver{}
	err := s.db.Get(&driver, "select * from drivers where id = $1", id)
	return driver, err
}

func (s *DriverService) GetByUserId(userId string) (model.Driver, error) {
	driver := model.Driver{}
	err := s.db.Get(&driver, "select * from drivers where user_id = $1", userId)
	return driver, err
}

func (s *DriverService) Create(userId string, driver model.DriverInput) (*model.Driver, error) {
	rows, err := s.db.Queryx(
		`insert into drivers (
        user_id,
        first_name,
        last_name,
        phone_number,
        policy_number,
        ins_provider,
        policy_start_date,
        policy_end_date,
        license_number,
        license_expire_date      
			)
	 	 values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	 	 returning *`,
		userId,
		driver.FirstName,
		driver.LastName,
		driver.PhoneNumber,
		driver.PolicyNumber,
		driver.InsProvider,
		driver.PolicyStartDate,
		driver.PolicyEndDate,
		driver.LicenseNumber,
		driver.LicenseExpireDate,
	)
	if err != nil {
		return nil, err
	}

	rtndriver := model.Driver{}

	rows.Next()
	err = rows.StructScan(&rtndriver)
	return &rtndriver, err
}
