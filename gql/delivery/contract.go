package delivery

import (
	"github.com/graphql-go/graphql"
	"github.com/yudhapratama10/search-gql/gql/usecase"
)

type resolver struct {
	footballUsecase usecase.FootballUsecaseContract
}

type Resolver interface {
	Search(params graphql.ResolveParams) (interface{}, error)
	Autocomplete(params graphql.ResolveParams) (interface{}, error)
}

func NewResolver(footballUsecase usecase.FootballUsecaseContract) Resolver {
	return &resolver{
		footballUsecase: footballUsecase,
	}
}
