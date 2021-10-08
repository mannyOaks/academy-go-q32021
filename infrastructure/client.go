package infrastructure

import (
	"encoding/json"
	"errors"
	"mrobles_app/common"

	"github.com/go-resty/resty/v2"
)

type MovieRepo struct{}

const baseUrl = "https://api.themoviedb.org/3"
const omdbAuthToken = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJhMGJkYWRhMmM5NTFhOTBiNmQxNjc4NjgyMTQ3NTRhMSIsInN1YiI6IjYxNWI5OTZjYzhhMmQ0MDAyYWMxMGM3YiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.u9kwuL1lNbkvWKUhPqP6vVssioOMiv7a94Wa3cmOm4E"

func newClient() *resty.Client {
	return resty.New().SetAuthToken(omdbAuthToken)
}

func (mr MovieRepo) GetMovie(id string) (common.Movie, error) {
	res, err := newClient().R().Get(baseUrl + "/movie/" + id)
	if err != nil {
		return common.Movie{}, err
	}

	if res.IsError() {
		return common.Movie{}, errors.New("Empty data")
	}
	return parseJsonMovie(res.Body())
}

func parseJsonMovie(body []byte) (common.Movie, error) {
	var data apiMovieResponse
	err := json.Unmarshal(body, &data)
	if err != nil {
		return common.Movie{}, err
	}

	movie := common.Movie{
		ID:          data.ID,
		Title:       data.Title,
		Overview:    data.Overview,
		Language:    data.OriginalLanguage,
		ReleaseDate: data.ReleaseDate,
		Poster:      data.PosterPath,
		Popularity:  data.Popularity,
		Adult:       data.Adult,
	}

	return movie, nil
}
