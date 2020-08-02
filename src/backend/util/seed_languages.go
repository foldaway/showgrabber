package util

import (
	"github.com/bottleneckco/showgrabber/src/backend/db"
	"github.com/bottleneckco/showgrabber/src/backend/model"
)

func SeedLanguages() error {
	var tvdbClient = CreateTVDBClient()
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
