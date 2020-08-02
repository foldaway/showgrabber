package util

import (
	"os"

	"github.com/bottleneckco/showgrabber/src/backend/db"
	"github.com/bottleneckco/showgrabber/src/backend/model"
	"github.com/pioz/tvdb"
)

var (
	TVDB_API_KEY   = os.Getenv("TVDB_API_KEY")
	TVDB_USER_KEY  = os.Getenv("TVDB_USER_KEY")
	TVDB_USER_NAME = os.Getenv("TVDB_USER_NAME")

	tvdbClient = tvdb.Client{
		Apikey:   TVDB_API_KEY,
		Userkey:  TVDB_USER_KEY,
		Username: TVDB_USER_NAME,
	}
)

func SeedLanguages() error {
	languages, err := tvdbClient.GetLanguages()
	if err != nil {
		return err
	}

	for _, lang := range languages {
		var dbLang model.Language
		err = db.DB.
			Where(model.Language{Abbreviation: lang.Abbreviation}).
			Assign(model.Language{
				Abbreviation: lang.Abbreviation,
				EnglishName:  lang.EnglishName,
				TVDBID:       lang.ID,
				Name:         lang.Name,
			}).
			FirstOrCreate(&dbLang).
			Error

		if err != nil {
			return err
		}
	}

	return nil
}
