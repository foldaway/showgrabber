package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/mrobinsn/go-newznab/newznab"
)

func (r *newznabResolver) Imdb(ctx context.Context, obj *newznab.NZB) (*string, error) {
	var id = obj.IMDBID
	return &id, nil
}

func (r *newznabResolver) Imdbscore(ctx context.Context, obj *newznab.NZB) (*float64, error) {
	var score = float64(obj.IMDBScore)
	return &score, nil
}

// Newznab returns generated.NewznabResolver implementation.
func (r *Resolver) Newznab() generated.NewznabResolver { return &newznabResolver{r} }

type newznabResolver struct{ *Resolver }
