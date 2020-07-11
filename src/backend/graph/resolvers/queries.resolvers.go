package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bottleneckco/showgrabber/src/backend/db"
	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/bottleneckco/showgrabber/src/backend/model"
	"github.com/pioz/tvdb"
)

func (r *queryResolver) Series(ctx context.Context) ([]*model.Series, error) {
	var results []*model.Series
	var err = db.DB.Model(&model.Series{}).Find(&results).Error
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
