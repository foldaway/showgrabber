// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/bottleneckco/showgrabber/src/backend/model"
)

type ParsedMetadata struct {
	SeasonNumber  *int    `json:"season_number"`
	EpisodeNumber *int    `json:"episode_number"`
	VideoCodec    *string `json:"video_codec"`
	AudioCodec    *string `json:"audio_codec"`
	Resolution    *string `json:"resolution"`
	SceneName     *string `json:"scene_name"`
	ReleaseFormat *string `json:"release_format"`
}

type SeriesAddInput struct {
	TvdbID int `json:"tvdbID"`
}

type SeriesAddPayload struct {
	Ok     bool          `json:"ok"`
	Series *model.Series `json:"series"`
}

type SeriesUpdateLanguageInput struct {
	SeriesID   int `json:"seriesID"`
	LanguageID int `json:"languageID"`
}

type SeriesUpdateLanguagePayload struct {
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
