package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/bottleneckco/showgrabber/src/backend/db"
	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/bottleneckco/showgrabber/src/backend/graph/model"
	"github.com/pioz/tvdb"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
)

func (r *mutationResolver) SeriesAdd(ctx context.Context, input model.SeriesAddInput) (*model.SeriesAddPayload, error) {
	var result model.SeriesAddPayload

	var tvdbSeries = &tvdb.Series{
		ID: input.TvdbID,
	}
	var err = tvdbClient.GetSeries(tvdbSeries)
	if err != nil {
		return nil, err
	}
	err = tvdbClient.GetSeriesPosterImages(tvdbSeries)
	if err != nil {
		log.Println(err)
	}

	var posterURL string
	if len(tvdbSeries.Images) > 0 {
		posterURL = tvdbSeries.Images[0].Thumbnail
	}

	var series = model.Series{
		Name:     tvdbSeries.SeriesName,
		Status:   tvdbSeries.Status,
		TvdbID:   input.TvdbID,
		Network:  tvdbSeries.Network,
		Poster:   posterURL,
		Overview: tvdbSeries.Overview,
	}

	err = db.DB.Create(&series).Error
	if err != nil {
		return nil, err
	}

	err = tvdbClient.GetSeriesSummary(tvdbSeries)
	if err != nil {
		return nil, err
	}

	var seasons []interface{}
	var seasonMap = make(map[int]model.Season)
	for _, seasonName := range tvdbSeries.Summary.AiredSeasons {
		seasonNum, err := strconv.Atoi(seasonName)
		if err != nil {
			return nil, err
		}
		var season = model.Season{
			Number:   uint(seasonNum),
			SeriesID: series.ID,
		}

		// NOTE: Need to insert one by one to get return ID
		err = db.DB.Create(&season).Error
		if err != nil {
			return nil, err
		}
		seasons = append(seasons, season)
		seasonMap[seasonNum] = season
	}

	// Fetch eps
	err = tvdbClient.GetSeriesEpisodes(tvdbSeries, url.Values{})
	if err != nil {
		return nil, err
	}

	var episodes []interface{}
	for _, episode := range tvdbSeries.Episodes {
		season, isSeasonExists := seasonMap[episode.AiredSeason]
		if !isSeasonExists {
			return nil, fmt.Errorf("Episode does not have a corresponding season: Title: %s, Ep #%d, Claimed Season: %d", episode.EpisodeName, episode.AiredEpisodeNumber, episode.AiredSeason)
		}
		airDate, err := time.Parse("2006-01-02", episode.FirstAired)
		if err != nil {
			log.Println(err)
			continue
		}
		var episode = model.Episode{
			Title:    episode.EpisodeName,
			Number:   episode.AiredEpisodeNumber,
			AirDate:  airDate,
			SeasonID: season.ID,
		}

		episodes = append(episodes, episode)
	}

	// NOTE: No need to return one by one, we don't care about episode ID
	err = gormbulk.BulkInsert(db.DB, episodes, 3000)
	if err != nil {
		return nil, err
	}

	result.Ok = err == nil
	result.Series = &series

	return &result, err
}

func (r *mutationResolver) SeriesUpdateLanguage(ctx context.Context, input model.SeriesUpdateLanguageInput) (*model.SeriesUpdateLanguagePayload, error) {
	var language model.Language
	var err = db.DB.First(&language, input.LanguageID).Error
	if err != nil {
		return nil, err
	}

	var series model.Series
	err = db.DB.First(&series, input.SeriesID).Error

	if err != nil {
		return nil, err
	}

	err = db.DB.Model(&series).Updates(model.Series{LanguageID: language.ID}).Error
	if err != nil {
		return nil, err
	}

	return &model.SeriesUpdateLanguagePayload{Ok: true, Series: &series}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
