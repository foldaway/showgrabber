package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/bottleneckco/showgrabber/src/backend/graph/generated"
	"github.com/bottleneckco/showgrabber/src/backend/graph/model"
	"github.com/bottleneckco/showgrabber/src/backend/util"
)

func (r *newznabResolver) ID(ctx context.Context, obj *model.Newznab) (*string, error) {
	if len(obj.ID) > 0 {
		return &obj.ID, nil
	}

	return &obj.Title, nil
}

func (r *newznabResolver) Imdb(ctx context.Context, obj *model.Newznab) (*string, error) {
	var id = obj.IMDBID
	return &id, nil
}

func (r *newznabResolver) Imdbscore(ctx context.Context, obj *model.Newznab) (*float64, error) {
	var score = float64(obj.IMDBScore)
	return &score, nil
}

func (r *newznabResolver) Parsed(ctx context.Context, obj *model.Newznab) (*model.ParsedMetadata, error) {
	return util.ParseNewznab(obj.Title)
}

// Newznab returns generated.NewznabResolver implementation.
func (r *Resolver) Newznab() generated.NewznabResolver { return &newznabResolver{r} }

type newznabResolver struct{ *Resolver }
