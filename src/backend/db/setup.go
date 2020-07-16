package db

import (
	"log"
	"os"

	"github.com/bottleneckco/showgrabber/src/backend/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var (
	dbURL = os.Getenv("DATABASE_URL")

	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("postgres", dbURL)

	if err != nil {
		log.Panic("Error connecting to database", err)
	}

	DB.AutoMigrate(
		&model.Series{},
		&model.Season{},
		&model.Episode{},
		&model.EpisodeFile{},
		&model.FetchJob{},
	)

	// Custom setup
	DB.Model(&model.Season{}).AddForeignKey("series_id", "series(id)", "CASCADE", "RESTRICT")
	DB.Model(&model.Episode{}).AddForeignKey("season_id", "seasons(id)", "CASCADE", "RESTRICT")
}
