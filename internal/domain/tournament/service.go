package tournament

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

func (s *Service) Collect() error {
	c := colly.NewCollector(
		/*colly.URLFilters(
			regexp.MustCompile("https://chess-results\\.com/t.+"),
		),*/
		colly.MaxDepth(2),
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		//fmt.Println(fmt.Sprintf("Checking '%s'", e.Attr("href")))
		if strings.HasPrefix(e.Attr("href"), "tnr") && strings.HasSuffix(e.Attr("href"), "lan=8") {
			tId := strings.Split(strings.TrimLeft(e.Attr("href"), "tnr"), ".")[0]
			s.repository.Save(tId, e.Text, e.Attr("href"))
			e.Request.Visit(e.Attr("href"))
		}
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://chess-results.com/fed.aspx?lan=8&fed=TUR")
	return nil
}
