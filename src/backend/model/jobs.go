package model

type FetchJob struct {
	BaseModel
	SeriesID  uint
	Series    Series
	EpisodeID uint
	Episode   Episode
}
