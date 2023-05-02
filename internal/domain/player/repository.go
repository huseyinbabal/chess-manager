package player

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	err := db.AutoMigrate(&Player{})
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}

func (r *Repository) Save(name, surname, fideID, federation string) (*Player, error) {
	e := Player{
		Name:       name,
		Surname:    surname,
		FideID:     fideID,
		Federation: federation,
	}
	tx := r.db.Save(&e)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &e, nil
}
