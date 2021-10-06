package services

import (
	"fmt"
	"mrobles_app/common"
	"mrobles_app/infrastructure"
)

const BASE_URL = "https://api.themoviedb.org/3/"

func fetchMovies() ([]common.Movie, error) {
	res, err := infrastructure.NewClient().R().Get(BASE_URL + "/discover/movie")
	if err != nil {
		return nil, err
	}

	movies, err := common.JsonToMovies(res.Body())
	if err != nil {
		return nil, err
	}

	fmt.Println(movies)

	infrastructure.Save(movies)
	return movies, nil
}

// FindMovies - Returns all the movie structs from the csv file and an error if one occurs
func FindMovies() ([]common.Movie, error) {
	return fetchMovies()
}

// FindMovie - Returns a movie struct from the csv file with specified id, if not found returns default value struct, and an error if one occurs
func FindMovie(id int) (common.Movie, error) {
	return infrastructure.FindOne(id)
}
