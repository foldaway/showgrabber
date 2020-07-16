// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/bottleneckco/showgrabber/src/backend/model"
)

type NewznabComment struct {
	Title   *string    `json:"title"`
	Content *string    `json:"content"`
	PubDate *time.Time `json:"pub_date"`
}

type SeriesAddInput struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Banner string `json:"banner"`
	TvdbID int    `json:"tvdbID"`
}

type SeriesAddPayload struct {
	Ok     bool          `json:"ok"`
	Series *model.Series `json:"series"`
}

type NewznabCategory string

const (
	NewznabCategoryTvAll        NewznabCategory = "TV_ALL"
	NewznabCategoryTvForeign    NewznabCategory = "TV_FOREIGN"
	NewznabCategoryTvSd         NewznabCategory = "TV_SD"
	NewznabCategoryTvHd         NewznabCategory = "TV_HD"
	NewznabCategoryTvUhd        NewznabCategory = "TV_UHD"
	NewznabCategoryTvOther      NewznabCategory = "TV_OTHER"
	NewznabCategoryTvSport      NewznabCategory = "TV_SPORT"
	NewznabCategoryMovieAll     NewznabCategory = "MOVIE_ALL"
	NewznabCategoryMovieForeign NewznabCategory = "MOVIE_FOREIGN"
	NewznabCategoryMovieOther   NewznabCategory = "MOVIE_OTHER"
	NewznabCategoryMovieSd      NewznabCategory = "MOVIE_SD"
	NewznabCategoryMovieHd      NewznabCategory = "MOVIE_HD"
	NewznabCategoryMovieUhd     NewznabCategory = "MOVIE_UHD"
	NewznabCategoryMovieBluray  NewznabCategory = "MOVIE_BLURAY"
	NewznabCategoryMovie3d      NewznabCategory = "MOVIE_3D"
)

var AllNewznabCategory = []NewznabCategory{
	NewznabCategoryTvAll,
	NewznabCategoryTvForeign,
	NewznabCategoryTvSd,
	NewznabCategoryTvHd,
	NewznabCategoryTvUhd,
	NewznabCategoryTvOther,
	NewznabCategoryTvSport,
	NewznabCategoryMovieAll,
	NewznabCategoryMovieForeign,
	NewznabCategoryMovieOther,
	NewznabCategoryMovieSd,
	NewznabCategoryMovieHd,
	NewznabCategoryMovieUhd,
	NewznabCategoryMovieBluray,
	NewznabCategoryMovie3d,
}

func (e NewznabCategory) IsValid() bool {
	switch e {
	case NewznabCategoryTvAll, NewznabCategoryTvForeign, NewznabCategoryTvSd, NewznabCategoryTvHd, NewznabCategoryTvUhd, NewznabCategoryTvOther, NewznabCategoryTvSport, NewznabCategoryMovieAll, NewznabCategoryMovieForeign, NewznabCategoryMovieOther, NewznabCategoryMovieSd, NewznabCategoryMovieHd, NewznabCategoryMovieUhd, NewznabCategoryMovieBluray, NewznabCategoryMovie3d:
		return true
	}
	return false
}

func (e NewznabCategory) String() string {
	return string(e)
}

func (e *NewznabCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NewznabCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NewznabCategory", str)
	}
	return nil
}

func (e NewznabCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
