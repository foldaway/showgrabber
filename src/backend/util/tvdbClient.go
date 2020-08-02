package util

import (
	"os"

	"github.com/pioz/tvdb"
)

var (
	TVDB_API_KEY   = os.Getenv("TVDB_API_KEY")
	TVDB_USER_KEY  = os.Getenv("TVDB_USER_KEY")
	TVDB_USER_NAME = os.Getenv("TVDB_USER_NAME")
)

func CreateTVDBClient() tvdb.Client {
	return tvdb.Client{
		Apikey:   TVDB_API_KEY,
		Userkey:  TVDB_USER_KEY,
		Username: TVDB_USER_NAME,
	}
}
