package model

import "time"

// Series represents a television series
type Series struct {
	BaseModel
	Name   string `json:"name"`
	Status string `json:"status"`
	Banner string `json:"banner"`
}

// Episode represents an episode in a series
type Episode struct {
	BaseModel
	Title    string    `json:"title"`
	Number   int       `json:"number"`
	AirDate  time.Time `json:"airDate"`
	SeriesID uint
	Series   Series
}

type EpisodeFile struct {
	BaseModel
	EpisodeId int
	Episode   Episode
	FilePath  string `json:"file_path"`
}
