package service

import (
	"errors"

	"github.com/conley21p/AccelDatabase/Server/internal/model"
	"github.com/jmoiron/sqlx"
)

type TrailerService struct {
	db *sqlx.DB
}

func NewTrailerService(db *sqlx.DB) *TrailerService {
	return &TrailerService{
		db: db,
	}
}

func (s *TrailerService) GetById(trailerId string, userId string) (*model.Trailer, error) {
	var trailer model.Trailer
	err := s.db.Get(&trailer,
		`SELECT t.* FROM trailers t 
		JOIN hauler_trailers ht ON t.id = ht.trailer_id
		JOIN haulers h ON ht.hauler_id = h.id
		JOIN drivers d ON h.driver_id = d.id
		WHERE t.id = $1 AND d.user_id = $2`,
		trailerId,
		userId)
	if err != nil {
		return nil, err
	}

	// Fetch associated hauler IDs
	haulerIds, err := s.fetchHaulerIds(trailer.Id)
	if err != nil {
		return nil, err
	}
	trailer.HaulerIds = haulerIds

	return &trailer, nil
}

func (s *TrailerService) fetchHaulerIds(trailerId string) ([]string, error) {
	var haulerIds []string
	err := s.db.Select(&haulerIds,
		`SELECT hauler_id FROM hauler_trailers WHERE trailer_id = $1`,
		trailerId)
	return haulerIds, err
}

func (s *TrailerService) updateHaulerAssociations(tx *sqlx.Tx, trailerId string, haulerIds []string) error {
	// First remove all existing associations
	_, err := tx.Exec(
		`DELETE FROM hauler_trailers WHERE trailer_id = $1`,
		trailerId,
	)
	if err != nil {
		return err
	}

	// Add new associations
	for _, haulerId := range haulerIds {
		_, err = tx.Exec(
			`INSERT INTO hauler_trailers (hauler_id, trailer_id) 
			VALUES ($1, $2)`,
			haulerId,
			trailerId,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *TrailerService) Create(trailer model.TrailerInput) (*model.Trailer, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Queryx(
		`INSERT INTO trailers (
			type,
			length,
			width,
			capacity,
			created_at
		) VALUES ($1, $2, $3, $4, NOW())
		RETURNING *`,
		trailer.Type,
		trailer.Length,
		trailer.Width,
		trailer.Capacity,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.Trailer
	if !rows.Next() {
		return nil, errors.New("failed to create trailer")
	}
	if err := rows.StructScan(&result); err != nil {
		return nil, err
	}

	// Handle hauler associations
	if len(trailer.HaulerIds) > 0 {
		if err := s.updateHaulerAssociations(tx, result.Id, trailer.HaulerIds); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *TrailerService) Update(trailerId string, userId string, input model.TrailerInput) (*model.Trailer, error) {
	rows, err := s.db.Queryx(
		`UPDATE trailers t
		SET type = $1,
			length = $2,
			width = $3,
			capacity = $4,
			updated_at = NOW()
		FROM haulers h
		JOIN drivers d ON h.driver_id = d.id
		WHERE t.id = $5 
		AND t.hauler_id = h.id
		AND d.user_id = $6
		RETURNING t.*`,
		input.Type,
		input.Length,
		input.Width,
		input.Capacity,
		trailerId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trailer model.Trailer
	if !rows.Next() {
		return nil, errors.New("trailer not found")
	}
	if err := rows.StructScan(&trailer); err != nil {
		return nil, err
	}

	return &trailer, nil
}

func (s *TrailerService) Delete(trailerId string, userId string) (*model.Trailer, error) {
	rows, err := s.db.Queryx(
		`DELETE FROM trailers t
		USING haulers h
		JOIN drivers d ON h.driver_id = d.id
		WHERE t.id = $1 
		AND t.hauler_id = h.id
		AND d.user_id = $2
		RETURNING t.*`,
		trailerId,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trailer model.Trailer
	if !rows.Next() {
		return nil, errors.New("trailer not found")
	}
	if err := rows.StructScan(&trailer); err != nil {
		return nil, err
	}

	return &trailer, nil
}
