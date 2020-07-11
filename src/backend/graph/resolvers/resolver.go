package resolvers

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pioz/tvdb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var (
	TVDB_API_KEY   = os.Getenv("TVDB_API_KEY")
	TVDB_USER_KEY  = os.Getenv("TVDB_USER_KEY")
	TVDB_USER_NAME = os.Getenv("TVDB_USER_NAME")

	tvdbClient = tvdb.Client{
		Apikey:   TVDB_API_KEY,
		Userkey:  TVDB_USER_KEY,
		Username: TVDB_USER_NAME,
	}
)

func init() {
	var err = tvdbClient.Login()
	if err != nil {
		log.Panic(err)
	}
}

type Resolver struct{}
