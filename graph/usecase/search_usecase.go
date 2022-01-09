package usecase

import (
	"fmt"

	"github.com/yudhapratama10/search-gql/graph/model"
)

func (usecase *footballUsecase) Search(keyword string, isPreOrder bool, page, take int) ([]model.Product, error) {
	data, err := usecase.repo.Search(keyword, isPreOrder, page, take)
	if err != nil {
		return []model.Product{}, err
	}

	fmt.Println(data)

	return data, nil

}
