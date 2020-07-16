package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/bottleneckco/showgrabber/src/backend/graph/model"
	"github.com/mrobinsn/go-newznab/newznab"
)

func (r *newznabResolver) ID(ctx context.Context, obj *newznab.NZB) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *newznabResolver) Comments(ctx context.Context, obj *newznab.NZB) ([]*model.NewznabComment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *newznabResolver) Imdb(ctx context.Context, obj *newznab.NZB) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *newznabResolver) Imdbscore(ctx context.Context, obj *newznab.NZB) (*float64, error) {
	panic(fmt.Errorf("not implemented"))
}

// Newznab returns generated.NewznabResolver implementation.
func (r *Resolver) Newznab() generated.NewznabResolver { return &newznabResolver{r} }

type newznabResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *newznabResolver) IsTorrent(ctx context.Context, obj *newznab.NZB) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
