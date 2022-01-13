package repository

import (
	"net/http"

	"github.com/yudhapratama10/search-gql/gql/model"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type footballRepository struct {
	http HTTPClient
}

type FootballRepositoryContract interface {
	Search(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error)
	Autocomplete(keyword string) ([]model.FootballClub, error)
}

func NewFootballRepository(http HTTPClient) FootballRepositoryContract {
	return &footballRepository{http: http}
}
