package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"net/url"
	"strconv"

	"github.com/bottleneckco/showgrabber/src/backend/db"
	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/bottleneckco/showgrabber/src/backend/model"
	"github.com/pioz/tvdb"
)

func (r *episodeResolver) ID(ctx context.Context, obj *model.Episode) (int, error) {
	return int(obj.ID), nil
}

func (r *seasonResolver) ID(ctx context.Context, obj *model.Season) (int, error) {
	return int(obj.ID), nil
}

func (r *seasonResolver) Number(ctx context.Context, obj *model.Season) (int, error) {
	return int(obj.Number), nil
}

func (r *seasonResolver) Episodes(ctx context.Context, obj *model.Season) ([]*model.Episode, error) {
	var dbResults []*model.Episode
	var err = db.DB.Order("number DESC").Where(&model.Episode{SeasonID: obj.ID}).Find(&dbResults).Error

	return dbResults, err
}

func (r *seriesResolver) ID(ctx context.Context, obj *model.Series) (int, error) {
	return int(obj.ID), nil
}

func (r *seriesResolver) Seasons(ctx context.Context, obj *model.Series) ([]*model.Season, error) {
	var dbResults []*model.Season
	var err = db.DB.Order("number DESC").Where(&model.Season{SeriesID: obj.ID}).Find(&dbResults).Error

	return dbResults, err
}

func (r *tVDBEpisodeResolver) SiteRating(ctx context.Context, obj *tvdb.Episode) (float64, error) {
	return float64(obj.SiteRating), nil
}

func (r *tVDBSeriesResolver) SiteRating(ctx context.Context, obj *tvdb.Series) (float64, error) {
	return float64(obj.SiteRating), nil
}

func (r *tVDBSeriesResolver) FanArtImages(ctx context.Context, obj *tvdb.Series) ([]*tvdb.Image, error) {
	var err = tvdbClient.GetSeriesFanartImages(obj)

	var results []*tvdb.Image

	for _, image := range obj.Images {
		var result = image

		results = append(results, &result)
	}

	if err != nil {
		log.Println(err)
	}

	return results, nil
}

func (r *tVDBSeriesResolver) PosterImages(ctx context.Context, obj *tvdb.Series) ([]*tvdb.Image, error) {
	var err = tvdbClient.GetSeriesPosterImages(obj)

	var results []*tvdb.Image

	for _, image := range obj.Images {
		var result = image

		results = append(results, &result)
	}

	if err != nil {
		log.Println(err)
	}

	return results, nil
}

func (r *tVDBSeriesResolver) SeasonImages(ctx context.Context, obj *tvdb.Series) ([]*tvdb.Image, error) {
	var err = tvdbClient.GetSeriesSeasonImages(obj)

	var results []*tvdb.Image

	for _, image := range obj.Images {
		var result = image

		results = append(results, &result)
	}

	if err != nil {
		log.Println(err)
	}

	return results, nil
}

func (r *tVDBSeriesResolver) SeasonWideImages(ctx context.Context, obj *tvdb.Series) ([]*tvdb.Image, error) {
	var err = tvdbClient.GetSeriesSeasonwideImages(obj)

	var results []*tvdb.Image

	for _, image := range obj.Images {
		var result = image

		results = append(results, &result)
	}

	if err != nil {
		log.Println(err)
	}

	return results, nil
}

func (r *tVDBSeriesResolver) SeriesImages(ctx context.Context, obj *tvdb.Series) ([]*tvdb.Image, error) {
	var err = tvdbClient.GetSeriesSeriesImages(obj)

	var results []*tvdb.Image

	for _, image := range obj.Images {
		var result = image

		results = append(results, &result)
	}

	if err != nil {
		log.Println(err)
	}

	return results, nil
}

func (r *tVDBSeriesResolver) Summary(ctx context.Context, obj *tvdb.Series) (*tvdb.Summary, error) {
	err := tvdbClient.GetSeriesSummary(obj)
	if err != nil {
		return nil, err
	}

	return &obj.Summary, nil
}

func (r *tVDBSeriesResolver) Episodes(ctx context.Context, obj *tvdb.Series, season *int) ([]*tvdb.Episode, error) {
	var results []*tvdb.Episode

	var values = url.Values{}

	if season != nil && len(obj.GetSeasonEpisodes(*season)) == 0 {
		values["airedSeason"] = []string{strconv.Itoa(*season)}
	}
	err := tvdbClient.GetSeriesEpisodes(obj, values)

	if err != nil {
		return results, err
	}

	if season != nil {
		return obj.GetSeasonEpisodes(*season), nil
	}

	for _, episode := range obj.Episodes {
		var result = episode
		results = append(results, &result)
	}

	return results, nil
}

// Episode returns generated.EpisodeResolver implementation.
func (r *Resolver) Episode() generated.EpisodeResolver { return &episodeResolver{r} }

// Season returns generated.SeasonResolver implementation.
func (r *Resolver) Season() generated.SeasonResolver { return &seasonResolver{r} }

// Series returns generated.SeriesResolver implementation.
func (r *Resolver) Series() generated.SeriesResolver { return &seriesResolver{r} }

// TVDBEpisode returns generated.TVDBEpisodeResolver implementation.
func (r *Resolver) TVDBEpisode() generated.TVDBEpisodeResolver { return &tVDBEpisodeResolver{r} }

// TVDBSeries returns generated.TVDBSeriesResolver implementation.
func (r *Resolver) TVDBSeries() generated.TVDBSeriesResolver { return &tVDBSeriesResolver{r} }

type episodeResolver struct{ *Resolver }
type seasonResolver struct{ *Resolver }
type seriesResolver struct{ *Resolver }
type tVDBEpisodeResolver struct{ *Resolver }
type tVDBSeriesResolver struct{ *Resolver }
