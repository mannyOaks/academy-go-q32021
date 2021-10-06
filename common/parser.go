package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// JsonToMovies - Returns array of movies parsed from the array of bytes of an API response
func JsonToMovies(body []byte) ([]Movie, error) {
	var jsonData apiDiscoverResponse
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err)
		return make([]Movie, 0), err
	}

	var movies []Movie
	for _, mov := range jsonData.Results {
		movie := Movie{
			ID:          mov.ID,
			Title:       mov.Title,
			Overview:    mov.Overview,
			Language:    mov.OriginalLanguage,
			ReleaseDate: mov.ReleaseDate,
			Poster:      mov.PosterPath,
			Popularity:  mov.Popularity,
			Adult:       mov.Adult,
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

// MoviesToStr - Returns array of rows of strings parsed from a movie array to be saved into the csv file
func MoviesToStr(movies []Movie) [][]string {
	var strs [][]string
	for _, mov := range movies {

		slice := []string{
			strconv.Itoa(mov.ID),
			mov.Title,
			mov.Overview,
			mov.Language,
			mov.Poster,
			strconv.FormatFloat(mov.Popularity, 'E', -1, 64),
			mov.ReleaseDate,
			strconv.FormatBool(mov.Adult),
		}
		strs = append(strs, slice)
	}

	return strs
}
