package service

import (
	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type InsuranceService struct {
	db *sqlx.DB
}

func NewInsuranceService(db *sqlx.DB) *InsuranceService {
	return &InsuranceService{
		db: db,
	}
}
func (s *InsuranceService) GetByDriverId(driverId string, userId string) (model.Insurance, error) {
	insurance := model.Insurance{}
	err := s.db.Get(&insurance,
		"select c.* from insurance i join drivers d on i.id = d.insurance_id where i.driver_id = $1 and d.user_id = $2",
		driverId,
		userId)
	if err != nil {
		return insurance, err
	}

	return insurance, nil
}

func (s *InsuranceService) Create(driverId string, insurance model.InsuranceInput) (*model.Insurance, error) {
	rows, err := s.db.Queryx(
		`INSERT INTO contact_info (
            driver_id, policy_number, ins_provider, policy_start_date, policy_end_date
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id, created_at, updated_at`,
		driverId,
		insurance.PolicyNumber,
		insurance.InsProvider,
		insurance.PolicyStartDate,
		insurance.PolicyEndDate,
	)
	if err != nil {
		return nil, err
	}

	rtncontact := model.Insurance{}

	rows.Next()
	err = rows.StructScan(&rtncontact)

	if err != nil {
		return nil, err
	}

	return &rtncontact, err
}
