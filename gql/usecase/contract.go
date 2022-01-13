package usecase

import (
	"github.com/yudhapratama10/search-gql/gql/model"
	"github.com/yudhapratama10/search-gql/gql/repository"
)

type footballUsecase struct {
	repo repository.FootballRepositoryContract
}

type FootballUsecaseContract interface {
	Search(keyword string, isPreOrder bool, page, take int) ([]model.FootballClub, error)
	Autocomplete(keyword string) ([]model.FootballClub, error)
}

func NewFootballUsecase(repo repository.FootballRepositoryContract) FootballUsecaseContract {
	return &footballUsecase{repo: repo}
}
