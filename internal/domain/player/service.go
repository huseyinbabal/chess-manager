package player

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Service struct {
	repository *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) Collect(tournamentUrl string) error {
	c := colly.NewCollector(
		/*colly.URLFilters(
			regexp.MustCompile("https://chess-results\\.com/t.+"),
		),*/
		colly.MaxDepth(2),
	)
	c.OnHTML(".TUR", func(e *colly.HTMLElement) {
		nameParts := strings.Split(e.ChildText("td:nth-of-type(3)"), ", ")
		fideID := e.ChildText("td:nth-of-type(4)")
		federation := e.ChildText("td:nth-of-type(5)")
		s.repository.Save(nameParts[1], nameParts[0], fideID, federation)
		//e.Request.Visit(e.Attr("href"))
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(tournamentUrl)
	return nil
}
