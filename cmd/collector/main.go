package main

import (
	"github.com/huseyinbabal/chess-manager/internal/config"
	"github.com/huseyinbabal/chess-manager/internal/domain/player"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	cfg := config.Load(os.Getenv("CONFIG_LOCATION"))
	db, err := gorm.Open(postgres.Open(cfg.DB.Dsn()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//tournamentRepo, err := tournament.NewRepository(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//tournamentSvc := tournament.NewService(tournamentRepo)

	playerRepo, err := player.NewRepository(db)

	if err != nil {
		log.Fatal(err)
	}
	playerSvc := player.NewService(playerRepo)

	playerSvc.Collect("https://chess-results.com/tnr759945.aspx?lan=8")

	/*s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minute().Do(func() {
		tournamentSvc.Collect()
	})
	s.StartBlocking()*/

}
