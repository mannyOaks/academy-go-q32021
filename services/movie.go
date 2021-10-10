package services

import (
	"strconv"

	"github.com/mannyOaks/academy-go-q32021/common"
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
	return MovieService{repo: repo}
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

	err := common.WriteToCsv(row)
	if err != nil {
		return err
	}

	return nil
}
