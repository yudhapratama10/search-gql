package repository

import (
	"github.com/yudhapratama10/search-gql/graph/model"
)

type footballRepository struct {
}

type FootballRepositoryContract interface {
	Search(keyword string, hasStadium bool, page, take int) ([]model.Product, error)
}

func NewFootballRepository() FootballRepositoryContract {
	return &footballRepository{}
}
