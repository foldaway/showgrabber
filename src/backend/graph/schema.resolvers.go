package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/bottleneckco/showgrabber/src/backend/graph/model"
	"github.com/pioz/tvdb"
)

func (r *queryResolver) Series(ctx context.Context) ([]*model.Series, error) {
	panic(fmt.Errorf("not implemented"))
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// TVDBSeries returns generated.TVDBSeriesResolver implementation.
func (r *Resolver) TVDBSeries() generated.TVDBSeriesResolver { return &tVDBSeriesResolver{r} }

type queryResolver struct{ *Resolver }
type tVDBSeriesResolver struct{ *Resolver }
