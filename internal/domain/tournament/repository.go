package tournament

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	err := db.AutoMigrate(&Tournament{})
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}

func (r *Repository) Save(id, name, url string) (*Tournament, error) {
	e := Tournament{
		Name: name,
		URL:  url,
	}
	e.ID = id
	tx := r.db.Save(&e)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &e, nil
}
