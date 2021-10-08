package services

import (
	"errors"
	"mrobles_app/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

var movie = common.Movie{
	ID:          635302,
	Title:       "Demon Slayer -Kimetsu no Yaiba- The Movie: Mugen Train",
	Overview:    "Tanjirō Kamado, joined with Inosuke Hashibira, a boy raised by boars who wears a boar's head, and Zenitsu Agatsuma, a scared boy who reveals his true power when he sleeps, boards the Infinity Train on a new mission with the Fire Hashira, Kyōjurō Rengoku, to defeat a demon who has been tormenting the people and killing the demon slayers who oppose it!",
	Language:    "ja",
	ReleaseDate: "2020-10-16",
	Poster:      "/h8Rb9gBr48ODIwYUttZNYeMWeUU.jpg",
	Popularity:  756.399,
	Adult:       false,
}

func TestMovieService_FindMovie(t *testing.T) {
	testCases := []struct {
		name     string
		response common.Movie
		hasError bool
		err      error
		id       string
	}{
		{
			name:     "id property",
			response: movie,
			err:      nil,
			hasError: false,
			id:       "635302",
		},
		{
			name:     "movie not found",
			response: common.Movie{},
			hasError: true,
			err:      errors.New("Movie 1 not found"),
			id:       "1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mockRepository{}
			mock.On("GetMovie", tc.id).Return(tc.response, tc.err)
			service := NewMovieService(&mock)

			movie, err := service.FindMovie(tc.id)
			assert.EqualValues(t, tc.response, movie)

			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}

}
