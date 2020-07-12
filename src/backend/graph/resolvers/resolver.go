package resolvers

import (
	"log"
	"os"

	graphModel "github.com/bottleneckco/showgrabber/src/backend/graph/model"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mrobinsn/go-newznab/newznab"
	"github.com/pioz/tvdb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var (
	TVDB_API_KEY     = os.Getenv("TVDB_API_KEY")
	TVDB_USER_KEY    = os.Getenv("TVDB_USER_KEY")
	TVDB_USER_NAME   = os.Getenv("TVDB_USER_NAME")
	NEWZNAB_BASE_URL = os.Getenv("NEWZNAB_BASE_URL")
	NEWZNAB_API_KEY  = os.Getenv("NEWZNAB_API_KEY")

	tvdbClient = tvdb.Client{
		Apikey:   TVDB_API_KEY,
		Userkey:  TVDB_USER_KEY,
		Username: TVDB_USER_NAME,
	}

	newznabClient = newznab.New(NEWZNAB_BASE_URL, NEWZNAB_API_KEY, 1, true)
)

func init() {
	var err = tvdbClient.Login()
	if err != nil {
		log.Panic(err)
	}
}

type Resolver struct{}

func convertNZBCategories(src []*graphModel.NewznabCategory) []int {
	var results []int

	for _, cat := range src {
		var result int
		switch *cat {
		case graphModel.NewznabCategoryTvAll:
			{
				result = newznab.CategoryTVAll
				break
			}
		case graphModel.NewznabCategoryTvForeign:
			{
				result = newznab.CategoryTVForeign
				break
			}
		case graphModel.NewznabCategoryTvSd:
			{
				result = newznab.CategoryTVSD
				break
			}
		case graphModel.NewznabCategoryTvHd:
			{
				result = newznab.CategoryTVHD
				break
			}
		case graphModel.NewznabCategoryTvUhd:
			{
				result = newznab.CategoryTVUHD
				break
			}
		case graphModel.NewznabCategoryTvOther:
			{
				result = newznab.CategoryTVOther
				break
			}
		case graphModel.NewznabCategoryTvSport:
			{
				result = newznab.CategoryTVSport
				break
			}
		case graphModel.NewznabCategoryMovieAll:
			{
				result = newznab.CategoryMovieAll
				break
			}
		case graphModel.NewznabCategoryMovieForeign:
			{
				result = newznab.CategoryMovieForeign
				break
			}
		case graphModel.NewznabCategoryMovieOther:
			{
				result = newznab.CategoryMovieOther
				break
			}
		case graphModel.NewznabCategoryMovieSd:
			{
				result = newznab.CategoryMovieSD
				break
			}
		case graphModel.NewznabCategoryMovieHd:
			{
				result = newznab.CategoryMovieHD
				break
			}
		case graphModel.NewznabCategoryMovieUhd:
			{
				result = newznab.CategoryMovieUHD
				break
			}
		case graphModel.NewznabCategoryMovieBluray:
			{
				result = newznab.CategoryMovieBluRay
				break
			}
		case graphModel.NewznabCategoryMovie3d:
			{
				result = newznab.CategoryMovie3D
				break
			}
		}

		results = append(results, result)
	}

	return results
}
