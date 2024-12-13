package service

import (
	"time"

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

func (s *DriverService) Create(driver model.Driver) (*model.Driver, error) {
	// Load the CST location
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Queryx(
		`insert into drivers (
			UserId            
			FirstName         
			LastName          
			PhoneNumber       
			PolicyNumber      
			InsProvider       
			PolicyStartDate   
			PolicyEndDate     
			LicenseNumber
			LicenseExpireDate 
			CreatedAt         
			)
	 	 values ($1, $2, $3, $4, $5, $6, $7)
	 	 returning *`,
		driver.UserId,
		driver.FirstName,
		driver.LastName,
		driver.PhoneNumber,
		driver.PolicyNumber,
		driver.InsProvider,
		driver.PolicyStartDate,
		driver.PolicyEndDate,
		driver.LicenseNumber,
		driver.LicenseExpireDate,
		time.Now().In(loc).Format("2006-01-02 15:04:05 CST"),
	)
	if err != nil {
		return nil, err
	}

	rtndriver := model.Driver{}

	rows.Next()
	err = rows.StructScan(&rtndriver)
	return &rtndriver, err
}
