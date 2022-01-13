package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yudhapratama10/search-gql/gql/model"
	"github.com/yudhapratama10/search-gql/gql/repository"
)

func TestAutocomplete(t *testing.T) {

	t.Parallel()

	t.Run("Should Success Autocomplete", func(t *testing.T) {
		var (
			data = []model.FootballClub{
				{
					Id:          25,
					Name:        "Newcastle United",
					Rating:      3.2,
					Description: "Newcastle United Football Club adalah klub sepak bola profesional Inggris yang berbasis di Newcastle upon Tyne, dan bermain di Liga Utama Inggris, kompetisi tingkat teratas dalam sepak bola Inggris. Newcastle United didirikan pada tahun 1892 sebagai hasil penggabungan Newcastle East End dan Newcastle West End, dan bermain di kandangnya saat ini, St James' Park, sejak saat itu. Stadion tersebut dikembangkan menjadi stadion all-seater pada pertengahan 1990-an dan memiliki kapasitas 52.354.",
				},
			}
			keyword = "Newc"
		)

		repo := new(repository.FootaballMock)
		uc := NewFootballUsecase(repo)

		repo.On("Autocomplete", keyword).Return(data, nil)

		resp, err := uc.Autocomplete(keyword)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
	})

	t.Run("Should Success Autocomplete 2", func(t *testing.T) {
		var (
			data = []model.FootballClub{
				{
					Id:          1,
					Name:        "FC Barcelona",
					Rating:      5,
					Description: "FC Barcelona adalah klub asal Spanyol.",
				},
			}
			keyword = "FC B"
		)

		repo := new(repository.FootaballMock)
		uc := NewFootballUsecase(repo)

		repo.On("Autocomplete", keyword).Return(data, nil)

		resp, err := uc.Autocomplete(keyword)
		assert.NoError(t, err)
		assert.NotEmpty(t, resp)
	})

}
