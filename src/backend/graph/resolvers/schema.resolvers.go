package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"net/url"
	"strconv"

	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/bottleneckco/showgrabber/src/backend/model"
	"github.com/pioz/tvdb"
)

func (r *seriesResolver) ID(ctx context.Context, obj *model.Series) (string, error) {
	return strconv.Itoa(int(obj.ID)), nil
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

func (r *tVDBSeriesResolver) Episodes(ctx context.Context, obj *tvdb.Series, season int) ([]*tvdb.Episode, error) {
	var results []*tvdb.Episode

	if len(obj.GetSeasonEpisodes(season)) == 0 {
		err := tvdbClient.GetSeriesEpisodes(obj, url.Values{
			"airedSeason": {strconv.Itoa(season)},
		})

		if err != nil {
			return results, err
		}
	}

	return obj.GetSeasonEpisodes(season), nil
}

// Series returns generated.SeriesResolver implementation.
func (r *Resolver) Series() generated.SeriesResolver { return &seriesResolver{r} }

// TVDBEpisode returns generated.TVDBEpisodeResolver implementation.
func (r *Resolver) TVDBEpisode() generated.TVDBEpisodeResolver { return &tVDBEpisodeResolver{r} }

// TVDBSeries returns generated.TVDBSeriesResolver implementation.
func (r *Resolver) TVDBSeries() generated.TVDBSeriesResolver { return &tVDBSeriesResolver{r} }

type seriesResolver struct{ *Resolver }
type tVDBEpisodeResolver struct{ *Resolver }
type tVDBSeriesResolver struct{ *Resolver }
