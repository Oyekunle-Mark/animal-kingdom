package store

import (
	"database/sql"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

type Store interface {
	CreateBird(bird *Bird) error
	GetBirds() ([]*Bird, error)
}

type DbStore struct {
	db *sql.DB
}

func (store *DbStore) CreateBird(bird *Bird) error {
	_, err := store.db.Query("INSERT INTO birds(species, description) VALUES ($1,$2)", bird.Species, bird.Description)
	return err
}

func (store *DbStore) GetBirds() ([]*Bird, error) {
	rows, err := store.db.Query("SELECT species, description from birds")

	if err != nil {
		return nil, err
	}
	
	defer rows.Close()

	birds := []*Bird{}

	for rows.Next() {
		bird := &Bird{}
		
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}
		
		birds = append(birds, bird)
	}
	return birds, nil
}

var store Store

func InitStore(s Store) {
	store = s
}