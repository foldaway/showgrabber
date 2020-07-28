package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

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

func (r *queryResolver) NzbSearchEpisode(ctx context.Context, categories []*graphModel.NewznabCategory, episodeID int) ([]*newznab.NZB, error) {
	var cat = categories
	if cat == nil {
		var tvAll = graphModel.NewznabCategoryTvAll
		cat = []*graphModel.NewznabCategory{&tvAll}
	}

	var ep model.Episode
	var results []*newznab.NZB

	var err = db.DB.First(&ep, episodeID).Error

	if err != nil {
		return results, err
	}

	err = db.DB.Model(&ep).Related(&ep.Season).Error
	if err != nil {
		return results, err
	}

	err = db.DB.Model(&ep.Season).Related(&ep.Season.Series).Error
	if err != nil {
		return results, err
	}

	var temp []newznab.NZB
	var newznabCats = convertNZBCategories(cat)
	var searchTerm = ""

	searchTerm = fmt.Sprintf("%s S%02dE%02d", ep.Season.Series.Name, ep.Season.Number, ep.Number)
	temp, err = newznabClient.SearchWithQuery(newznabCats, searchTerm, "search")
	log.Printf("[NEWZNAB] Search '%s', %d results, error: %+v\n", searchTerm, len(temp), err)

	for _, nzbResult := range temp {
		var result = nzbResult
		results = append(results, &result)
	}

	searchTerm = fmt.Sprintf("%s %s", ep.Season.Series.Name, ep.Title)
	temp, err = newznabClient.SearchWithQuery(newznabCats, searchTerm, "search")
	log.Printf("[NEWZNAB] Search '%s', %d results, error: %+v\n", searchTerm, len(temp), err)

	for _, nzbResult := range temp {
		var result = nzbResult
		results = append(results, &result)
	}

	return results, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
