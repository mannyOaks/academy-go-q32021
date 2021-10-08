package services

import (
	"mrobles_app/common"
	"strconv"
)

type repository interface {
	GetMovie(id string) (common.Movie, error)
}

type MovieService struct {
	repo repository
}

func NewMovieService(repo repository) MovieService {
	return MovieService{repo: repo}
}

const csvFilePath = "movies.csv"

// FindMovie - Returns and saves movie from api with the specified id
func (ms MovieService) FindMovie(id string) (common.Movie, error) {
	movie, err := ms.repo.GetMovie(id)
	if err != nil {
		return common.Movie{}, err
	}

	if err := saveToCsv(movie); err != nil {
		return common.Movie{}, err
	}
	return movie, nil
}

func saveToCsv(mov common.Movie) error {
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

	err := common.WriteToCsv(csvFilePath, row)
	if err != nil {
		return err
	}

	return nil
}
