package delivery

import (
	"github.com/graphql-go/graphql"
)

func (r *resolver) Search(params graphql.ResolveParams) (interface{}, error) {
	//ctx := context.Background()
	hasstadium := false
	keyword := ""
	page := 0
	take := 0

	if keywordFromClients, ok := params.Args["keyword"].(string); ok {
		keyword = keywordFromClients
	}

	if hasStadiumFromClients, ok := params.Args["hasstadium"].(bool); ok {
		hasstadium = hasStadiumFromClients
	}

	if pageFromClients, ok := params.Args["page"].(int); ok {
		page = pageFromClients
	}

	if takeFromClients, ok := params.Args["take"].(int); ok {
		take = takeFromClients
	}

	results, err := r.footballUsecase.Search(keyword, hasstadium, page, take)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r *resolver) Autocomplete(params graphql.ResolveParams) (interface{}, error) {
	//ctx := context.Background()
	keyword := ""
	// page := 0
	// take := 0

	if keywordFromClients, ok := params.Args["keyword"].(string); ok {
		keyword = keywordFromClients
	}

	results, err := r.footballUsecase.Autocomplete(keyword)
	if err != nil {
		return nil, err
	}

	return results, nil
}
