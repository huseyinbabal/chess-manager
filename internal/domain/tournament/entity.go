package tournament

import (
	"github.com/huseyinbabal/chess-manager/internal/gormx"
)

type Tournament struct {
	gormx.Model
	Name string
	URL  string
}
