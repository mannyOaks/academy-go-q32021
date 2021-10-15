package services

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/mannyOaks/academy-go-q32021/entities"
)

type movieRepository interface {
	GetMovie(id string) (entities.Movie, error)
}

// MovieService represents the service that will be used by a controller
type MovieService struct {
	repo movieRepository
}

// NewMovieService - receives a movieRepository and instantiates a movie service
func NewMovieService(repo movieRepository) MovieService {
	return MovieService{repo}
}

// FindMovie - Returns and saves movie from api with the specified id
func (ms MovieService) FindMovie(id string) (entities.Movie, error) {
	movie, err := ms.repo.GetMovie(id)
	if err != nil {
		return entities.Movie{}, err
	}

	if err := saveToCsv(movie); err != nil {
		return entities.Movie{}, err
	}
	return movie, nil
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
