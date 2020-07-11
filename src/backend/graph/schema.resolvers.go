package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"net/url"
	"strconv"

	"github.com/bottleneckco/showgrabber/src/backend/db"
	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/bottleneckco/showgrabber/src/backend/graph/model"
	mainModel "github.com/bottleneckco/showgrabber/src/backend/model"
	"github.com/pioz/tvdb"
)

func (r *queryResolver) Series(ctx context.Context) ([]*model.Series, error) {
	var results []*model.Series
	var err = db.DB.Model(&mainModel.Series{}).Find(&results).Error
	return results, err
}

func (r *queryResolver) TvdbSeriesSearch(ctx context.Context, term string) ([]*tvdb.Series, error) {
	var results []*tvdb.Series

	tvdbResults, err := tvdbClient.SearchByName(term)
	if err != nil {
		return results, err
	}

	for _, tvdbResult := range tvdbResults {
		var tempResult = tvdbResult
		results = append(results, &tempResult)
	}

	return results, nil
}

func (r *tVDBEpisodeResolver) SiteRating(ctx context.Context, obj *tvdb.Episode) (float64, error) {
	return float64(obj.SiteRating), nil
}

func (r *tVDBSeriesResolver) SiteRating(ctx context.Context, obj *tvdb.Series) (float64, error) {
	return float64(obj.SiteRating), nil
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TVDBEpisode returns generated.TVDBEpisodeResolver implementation.
func (r *Resolver) TVDBEpisode() generated.TVDBEpisodeResolver { return &tVDBEpisodeResolver{r} }

// TVDBSeries returns generated.TVDBSeriesResolver implementation.
func (r *Resolver) TVDBSeries() generated.TVDBSeriesResolver { return &tVDBSeriesResolver{r} }

type queryResolver struct{ *Resolver }
type tVDBEpisodeResolver struct{ *Resolver }
type tVDBSeriesResolver struct{ *Resolver }
