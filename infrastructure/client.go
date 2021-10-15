package infrastructure

import (
	"encoding/json"
	"errors"

	"github.com/mannyOaks/academy-go-q32021/entities"

	"github.com/go-resty/resty/v2"
)

type MovieRepo struct{}

func NewMovieRepo() MovieRepo {
	return MovieRepo{}
}

const baseUrl = "https://api.themoviedb.org/3"
const omdbAuthToken = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJhMGJkYWRhMmM5NTFhOTBiNmQxNjc4NjgyMTQ3NTRhMSIsInN1YiI6IjYxNWI5OTZjYzhhMmQ0MDAyYWMxMGM3YiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.u9kwuL1lNbkvWKUhPqP6vVssioOMiv7a94Wa3cmOm4E"

func newClient() *resty.Client {
	return resty.New().SetAuthToken(omdbAuthToken)
}

func (mr MovieRepo) GetMovie(id string) (entities.Movie, error) {
	res, err := newClient().R().Get(baseUrl + "/movie/" + id)
	if err != nil {
		return entities.Movie{}, err
	}

	if res.IsError() {
		return entities.Movie{}, errors.New("empty data")
	}
	return parseJsonMovie(res.Body())
}

func parseJsonMovie(body []byte) (entities.Movie, error) {
	var data entities.ApiMovieResponse
	err := json.Unmarshal(body, &data)
	if err != nil {
		return entities.Movie{}, err
	}

	movie := entities.Movie{
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
