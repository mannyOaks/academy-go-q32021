package services

import (
	"fmt"
	"math"
	"testing"

	"github.com/mannyOaks/academy-go-q32021/entities"
	"github.com/mannyOaks/academy-go-q32021/services/mocks"

	"github.com/stretchr/testify/assert"
)

var (
	movie = &entities.Movie{
		ID:          635302,
		Title:       "Demon Slayer -Kimetsu no Yaiba- The Movie: Mugen Train",
		Overview:    "Tanjirō Kamado, joined with Inosuke Hashibira, a boy raised by boars who wears a boar's head, and Zenitsu Agatsuma, a scared boy who reveals his true power when he sleeps, boards the Infinity Train on a new mission with the Fire Hashira, Kyōjurō Rengoku, to defeat a demon who has been tormenting the people and killing the demon slayers who oppose it!",
		Language:    "ja",
		ReleaseDate: "2020-10-16",
		Poster:      "/h8Rb9gBr48ODIwYUttZNYeMWeUU.jpg",
		Popularity:  756.399,
		Adult:       false,
	}
	oddMovies = []entities.Movie{
		{
			ID:          76341,
			Title:       "Mad Max: Fury Road",
			Overview:    "An apocalyptic story set in the furthest reaches of our planet, in a stark desert landscape where humanity is broken, and most everyone is crazed fighting for the necessities of life. Within this world exist two rebels on the run who just might be able to restore order.",
			Language:    "en",
			ReleaseDate: "2015-05-13",
			Poster:      "/8tZYtuWezp8JbcsvHYO0O46tFbo.jpg",
			Popularity:  7.6362,
			Adult:       false,
		},
		{
			ID:          703771,
			Title:       "Deathstroke: Knights & Dragons - The Movie",
			Overview:    "The assassin Deathstroke tries to save his family from the wrath of H.I.V.E. and the murderous Jackal.",
			Language:    "en",
			ReleaseDate: "2020-08-04",
			Poster:      "/vFIHbiy55smzi50RmF8LQjmpGcx.jpg",
			Popularity:  2.055471,
			Adult:       false,
		},
	}
	evenMovies = []entities.Movie{
		{
			ID:          635302,
			Title:       "Demon Slayer -Kimetsu no Yaiba- The Movie: Mugen Train",
			Overview:    "Tanjirō Kamado, joined with Inosuke Hashibira, a boy raised by boars who wears a boar's head, and Zenitsu Agatsuma, a scared boy who reveals his true power when he sleeps, boards the Infinity Train on a new mission with the Fire Hashira, Kyōjurō Rengoku, to defeat a demon who has been tormenting the people and killing the demon slayers who oppose it!",
			Language:    "ja",
			ReleaseDate: "2020-10-16",
			Poster:      "/h8Rb9gBr48ODIwYUttZNYeMWeUU.jpg",
			Popularity:  1.102931,
			Adult:       false,
		},
		{
			ID:          39254,
			Title:       "Real Steel",
			Overview:    "Charlie Kenton is a washed-up fighter who retired from the ring when robots took over the sport. After his robot is trashed, he reluctantly teams up with his estranged son to rebuild and train an unlikely contender.",
			Language:    "en",
			ReleaseDate: "2011-09-28",
			Poster:      "/4GIeI5K5YdDUkR3mNQBoScpSFEf.jpg",
			Popularity:  2.00198,
			Adult:       false,
		},
		{
			ID:          45682,
			Title:       "Jackie Chan Kung Fu Master",
			Overview:    "Jackie Chan is the undefeated Kung Fu Master who dishes out the action in traditional Jackie Chan style. When a young boy sets out to learn how to fight from the Master himself, he not only witnesses some spectacular fights, but learns some important life lessons along the way.",
			Language:    "zh",
			ReleaseDate: "2009-07-03",
			Poster:      "/ds8xP7319zuPMa09kxzkIPBsHVL.jpg",
			Popularity:  1.46423,
			Adult:       false,
		},
	}
)

func TestMovieService_FindMovie(t *testing.T) {
	testCases := []struct {
		name     string
		response *entities.Movie
		id       string
	}{
		{
			name:     "id property",
			response: movie,
			id:       "635302",
		},
		{
			name:     "movie not found",
			response: &entities.Movie{},
			id:       "1",
		},
	}

	t.Setenv("CSV_PATH", "movies.csv")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mockRepo := mocks.MovieRepository{}
			mockWPool := mocks.WorkerPool{}
			mockRepo.On("GetMovie", tc.id).Return(*tc.response, nil)
			service := NewMovieService(&mockRepo, &mockWPool)

			movie, err := service.FindMovie(tc.id)
			fmt.Println(movie, err)

			assert.Nil(t, err)
			assert.EqualValues(t, tc.response, movie)
		})
	}

}

func TestMovieService_FindMovies(t *testing.T) {
	testCases := []struct {
		name           string
		filter         string
		items          int
		itemsPerWorker int
		workersNum     int
		err            error
		value          []entities.Movie
	}{
		{
			name:           "get odd movies",
			filter:         "odd",
			items:          2,
			itemsPerWorker: 1,
			workersNum:     2,
			err:            nil,
			value:          oddMovies,
		},
		{
			name:           "get even movies",
			filter:         "odd",
			items:          3,
			itemsPerWorker: 2,
			workersNum:     2,
			err:            nil,
			value:          evenMovies,
		},
	}

	t.Setenv("CSV_PATH", "movies.csv")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tasks := int(math.Ceil(float64(tc.items) / float64(tc.itemsPerWorker)))

			mockRepo := mocks.MovieRepository{}
			mockWPool := mocks.WorkerPool{}
			mockWPool.On("GetMovies", tc.filter, tasks, tc.items, tc.itemsPerWorker).Return(tc.value, tc.err)
			service := NewMovieService(&mockRepo, &mockWPool)

			value, err := service.FindMovies(tc.filter, tc.items, tc.itemsPerWorker)
			workersNum := service.GetWorkers(tc.items, tc.itemsPerWorker)

			if tc.err != nil {
				assert.EqualError(t, tc.err, err.Error())
			} else {
				assert.Nil(t, err)
			}

			assert.EqualValues(t, tc.value, value)
			assert.EqualValues(t, len(tc.value), len(value))

			assert.Equal(t, tc.workersNum, workersNum)
		})
	}
}
