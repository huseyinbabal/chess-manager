package player

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name       string
	Surname    string
	FideID     string
	Federation string
}
