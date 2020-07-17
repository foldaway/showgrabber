package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bottleneckco/showgrabber/src/backend/db"
	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	graphModel "github.com/bottleneckco/showgrabber/src/backend/graph/model"
	"github.com/bottleneckco/showgrabber/src/backend/model"
	"github.com/mrobinsn/go-newznab/newznab"
	"github.com/pioz/tvdb"
)

func (r *queryResolver) Series(ctx context.Context) ([]*model.Series, error) {
	var results []*model.Series
	var err = db.DB.Model(&model.Series{}).Find(&results).Error
	return results, err
}

func (r *queryResolver) SeriesByID(ctx context.Context, id *int) (*model.Series, error) {
	var result model.Series
	var err = db.DB.Model(&model.Series{}).First(&result, *id).Error
	return &result, err
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

func (r *queryResolver) NzbSearch(ctx context.Context, categories []*graphModel.NewznabCategory, term string) ([]*newznab.NZB, error) {
	var results []*newznab.NZB

	nzbResults, err := newznabClient.SearchWithQuery(convertNZBCategories(categories), term, "search")

	for _, nzbResult := range nzbResults {
		var result = nzbResult
		results = append(results, &result)
	}

	return results, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
