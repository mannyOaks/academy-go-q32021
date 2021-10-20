package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/mannyOaks/academy-go-q32021/controllers/mocks"
	"github.com/mannyOaks/academy-go-q32021/entities"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	movieJson      = "{\"movie\":{\"id\":635302,\"title\":\"Demon Slayer -Kimetsu no Yaiba- The Movie: Mugen Train\",\"description\":\"Tanjirō Kamado, joined with Inosuke Hashibira, a boy raised by boars who wears a boar's head, and Zenitsu Agatsuma, a scared boy who reveals his true power when he sleeps, boards the Infinity Train on a new mission with the Fire Hashira, Kyōjurō Rengoku, to defeat a demon who has been tormenting the people and killing the demon slayers who oppose it!\",\"language\":\"ja\",\"release_date\":\"2020-10-16\",\"poster_path\":\"/h8Rb9gBr48ODIwYUttZNYeMWeUU.jpg\",\"popularity\":756.399,\"adult\":false}}\n"
	notFoundJson   = "{\"message\":\"Movie %s not found\"}\n"
	badRequestJson = "{\"message\":\"Param {id} must be numeric\"}\n"
	movie          = &entities.Movie{
		ID:          635302,
		Title:       "Demon Slayer -Kimetsu no Yaiba- The Movie: Mugen Train",
		Overview:    "Tanjirō Kamado, joined with Inosuke Hashibira, a boy raised by boars who wears a boar's head, and Zenitsu Agatsuma, a scared boy who reveals his true power when he sleeps, boards the Infinity Train on a new mission with the Fire Hashira, Kyōjurō Rengoku, to defeat a demon who has been tormenting the people and killing the demon slayers who oppose it!",
		Language:    "ja",
		ReleaseDate: "2020-10-16",
		Poster:      "/h8Rb9gBr48ODIwYUttZNYeMWeUU.jpg",
		Popularity:  756.399,
		Adult:       false,
	}
)

func TestMovieController_GetMovie(t *testing.T) {
	testCases := []struct {
		name     string
		response string
		hasError bool
		err      error
		id       string
		status   int
		movie    *entities.Movie
		path     string
		param    string
	}{
		{
			name:     "find kimetsu no yaiba",
			response: movieJson,
			err:      nil,
			id:       "85937",
			status:   http.StatusOK,
			movie:    movie,
			path:     "/movies/:id",
			param:    "id",
		},
		{
			name:     "not found",
			response: fmt.Sprintf(notFoundJson, "10"),
			err:      nil,
			id:       "10",
			status:   http.StatusNotFound,
			movie:    nil,
			path:     "/movies/:id",
			param:    "id",
		},
		{
			name:     "wrong param type",
			response: badRequestJson,
			err:      nil,
			id:       "askjdnaskldnalsndlasndklans",
			status:   http.StatusBadRequest,
			movie:    nil,
			path:     "/movies/:id",
			param:    "id",
		},
	}

	e := echo.New()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mocks.MovieService{}
			mock.On("FindMovie", tc.id).Return(tc.movie, tc.err)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath(tc.path)
			c.SetParamNames(tc.param)
			c.SetParamValues(tc.id)
			h := NewMovieController(&mock)

			res := h.GetMovie(c)
			if tc.err != nil {
				assert.EqualError(t, tc.err, res.Error())
			}

			assert.Equal(t, tc.status, rec.Code)
			assert.Equal(t, tc.response, rec.Body.String())
		})

	}

}

func TestMovieController_GetMovies(t *testing.T) {
	testCases := []struct {
		name           string
		filter         string
		items          int
		itemsPerWorker int
		err            error
		movies         []entities.Movie
		path           string
		response       string
	}{}

	e := echo.New()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := mocks.MovieService{}
			mock.On("FindMovies", tc.filter, tc.items, tc.itemsPerWorker).Return(tc.movies, tc.err)

			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			q := make(url.Values)
			q.Set("type", tc.filter)
			q.Set("items", strconv.Itoa(tc.items))
			q.Set("items_per_worker", strconv.Itoa(tc.itemsPerWorker))

			c.SetPath(tc.path + "/?" + q.Encode())
			h := NewMovieController(&mock)

			response := h.GetMovies(c)

			if tc.err == nil {
				assert.NoError(t, response)
			}

			assert.EqualValues(t, tc.response, response)
		})
	}
}
