package repository

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/yudhapratama10/search-gql/gql/model"
)

// Service Base URL
var baseURL string = "http://localhost:8080"

func (repo *footballRepository) Search(keyword string, hasStadium bool, page, take int) ([]model.Product, error) {
	var err error
	var client = &http.Client{}
	var res []model.Product

	// req, err := json.Marshal(model.ProductParams{
	// 	Keyword:    keyword,
	// 	HasStadium: hasStadium,
	// 	Page:       page,
	// 	Take:       take,
	// })

	// if err != nil {
	// 	return nil, err
	// }

	// request, err := http.NewRequest("GET", baseURL+"/search", bytes.NewBuffer(req))

	params := "keyword=" + url.QueryEscape(keyword) + "&hasstadium=" + strconv.FormatBool(hasStadium) + "&page=" + strconv.Itoa(page) + "&take=" + strconv.Itoa(take)

	request, err := http.NewRequest("GET", baseURL+"/search?"+params, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
