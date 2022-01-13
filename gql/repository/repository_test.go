package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yudhapratama10/search-gql/gql/model"
)

func TestAutocomplete(t *testing.T) {

	client := &HTTPClientMock{}

	var data = []model.FootballClub{
		{
			Id:          25,
			Name:        "Newcastle United",
			Rating:      2.5,
			Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
			HasStadium:  true,
			Tournaments: []string{"English Premier League"},
			// Nation:      "Inggris",
			// Tournaments: []string{"English Premier League", "FA Cup"},
		},
	}

	reqByte, _ := json.Marshal(data)
	r := io.NopCloser(strings.NewReader(string(reqByte)))

	// create a new reader with that JSON
	client.DoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	repo := NewFootballRepository(client)
	resp, err := repo.Autocomplete("string")

	// fmt.Println(resp)

	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

}
