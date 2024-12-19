package service

import (
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
func (s *ContactInfoService) GetByDriverId(driverId string, userId string) (model.ContactInfo, error) {
	contactInfo := model.ContactInfo{}
	err := s.db.Get(&contactInfo,
		"select c.* from contact_info c join drivers d on c.id = d.contact_info_id where c.driver_id = $1 and d.user_id = $2",
		driverId,
		userId)
	if err != nil {
		return contactInfo, err
	}

	return contactInfo, nil
}

func (s *ContactInfoService) Create(driverId string, contactInfo model.ContactInfoInput) (*model.ContactInfo, error) {
	rows, err := s.db.Queryx(
		`INSERT INTO contact_info (
            driver_id, street_address, city, state, zip_code, country
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, created_at, updated_at`,
		driverId,
		contactInfo.StreetAddress,
		contactInfo.City,
		contactInfo.State,
		contactInfo.ZipCode,
		contactInfo.Country,
	)
	if err != nil {
		return nil, err
	}

	rtncontact := model.ContactInfo{}

	rows.Next()
	err = rows.StructScan(&rtncontact)

	if err != nil {
		return nil, err
	}

	return &rtncontact, err
}
