package services

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"

	"github.com/mannyOaks/academy-go-q32021/entities"
)

type MovieRepository interface {
	GetMovie(id string) (entities.Movie, error)
}

type WorkerPool interface {
	GetMovies(string, int, int, int) ([]entities.Movie, error)
}

// MovieService represents the service that will be used by a controller
type MovieService struct {
	repo MovieRepository
	pool WorkerPool
}

// NewMovieService - receives a movieRepository and instantiates a movie service
func NewMovieService(repo MovieRepository, pool WorkerPool) MovieService {
	return MovieService{repo, pool}
}

// FindMovie - Returns and saves movie from api with the specified id
func (ms MovieService) FindMovie(id string) (*entities.Movie, error) {
	movie, err := ms.repo.GetMovie(id)
	if err != nil {
		return nil, err
	}

	if err := saveToCsv(movie); err != nil {
		return nil, err
	}
	return &movie, nil
}

// FindMovies - returns an array of movies read from a csv file and an error
func (ms MovieService) FindMovies(filter string, items int, itemsPerWorker int) ([]entities.Movie, error) {
	numWorkers := ms.GetWorkers(items, itemsPerWorker)
	movies, err := ms.pool.GetMovies(filter, numWorkers, items, itemsPerWorker)
	if err != nil {
		return nil, err
	}

	if len(movies) > items {
		movies = movies[:items]
	}
	return movies, nil
}

// GetWorkers- returns total of workers the pool should have
func (ms MovieService) GetWorkers(items, itemsPerWorker int) int {
	return int(math.Ceil(float64(items) / float64(itemsPerWorker)))
}

func saveToCsv(mov entities.Movie) error {
	// perms that work => 0644, 066, 0755
	file, err := os.OpenFile(os.Getenv("CSV_PATH"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)

	row := []string{
		strconv.Itoa(mov.ID),
		mov.Title,
		mov.Overview,
		mov.Language,
		mov.Poster,
		strconv.FormatFloat(mov.Popularity, 'E', -1, 64),
		mov.ReleaseDate,
		strconv.FormatBool(mov.Adult),
	}

	err = w.WriteAll([][]string{row})
	if err != nil {
		return err
	}

	return nil
}
