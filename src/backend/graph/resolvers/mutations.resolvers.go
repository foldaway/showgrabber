package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bottleneckco/showgrabber/src/backend/db"
	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/bottleneckco/showgrabber/src/backend/graph/model"
	"github.com/pioz/tvdb"
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

	var series = model.Series{
		Name:   tvdbSeries.SeriesName,
		Status: tvdbSeries.Status,
		TvdbID: input.TvdbID,
	}

	err = db.DB.Create(&series).Error

	result.Ok = err == nil
	result.Series = &series

	return &result, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
