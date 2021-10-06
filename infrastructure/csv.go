package infrastructure

import (
	"mrobles_app/common"
	"strconv"
)

const csvFilePath = "movies.csv"

func parseCsv() ([]common.Movie, error) {
	records, err := common.ReadCsvFile(csvFilePath)
	if err != nil {
		return make([]common.Movie, 0), err
	}

	slice := make([]common.Movie, len(records[0]))

	for _, rec := range records {
		id, _ := strconv.Atoi(rec[0])

		game := common.Movie{
			ID: id,
			// Name:        rec[1],
			ReleaseDate: rec[2],
		}

		slice = append(slice, game)
	}

	return slice, nil
}

// FindOne - Returns a movie with the specified id
func FindOne(id int) (common.Movie, error) {
	movies, err := parseCsv()
	if err != nil {
		return common.Movie{}, err
	}

	for _, mov := range movies {
		if mov.ID == id {
			return mov, nil
		}
	}

	return common.Movie{}, nil
}

// FindAll - Returns all the records in the csv file
func FindAll() ([]common.Movie, error) {
	return parseCsv()
}

// Save - Saves an array of movies to the csv file
func Save(data []common.Movie) {
	common.WriteToCsv(csvFilePath, common.MoviesToStr(data))
}
