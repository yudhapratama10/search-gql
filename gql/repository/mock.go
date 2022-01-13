package repository

import (
	"net/http"

	"github.com/stretchr/testify/mock"
	"github.com/yudhapratama10/search-gql/gql/model"
)

type HTTPClientMock struct {
	// DoFunc will be executed whenever Do function is executed
	// so we'll be able to create a custom response
	DoFunc func(*http.Request) (*http.Response, error)
}

func (h HTTPClientMock) Do(r *http.Request) (*http.Response, error) {
	return h.DoFunc(r)
}

//======================================================================

type FootaballMock struct {
	mock.Mock
}

func (f *FootaballMock) Search(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error) {
	args := f.Called(keyword, hasStadium, page, take)

	return args.Get(0).([]model.FootballClub), args.Error(1)
}

func (f *FootaballMock) Autocomplete(keyword string) ([]model.FootballClub, error) {
	args := f.Called(keyword)

	return args.Get(0).([]model.FootballClub), args.Error(1)
}
