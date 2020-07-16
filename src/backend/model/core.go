package model

import "time"

// Series represents a television series
type Series struct {
	BaseModel
	Name    string   `json:"name"`
	Status  string   `json:"status"`
	Network string   `json:"network"`
	Poster  string   `json:"poster"`
	TvdbID  int      `json:"tvdbID"`
	Seasons []Season `json:"seasons"`
}

// Season represents a season in a series
type Season struct {
	BaseModel
	Name     string    `json:"name"`
	Episodes []Episode `json:"episodes"`
	SeriesID uint
	Series   Series `json:"series"`
}

// Episode represents an episode in a series
type Episode struct {
	BaseModel
	Title    string    `json:"title"`
	Number   int       `json:"number"`
	AirDate  time.Time `json:"airDate"`
	SeasonID uint
	Season   Season `json:"season"`
}

type EpisodeFile struct {
	BaseModel
	EpisodeId int
	Episode   Episode
	FilePath  string `json:"file_path"`
}
