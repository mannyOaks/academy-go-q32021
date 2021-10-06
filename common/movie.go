package common

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Overview    string  `json:"description"`
	Language    string  `json:"language"`
	ReleaseDate string  `json:"release_date"`
	Poster      string  `json:"poster_path"`
	Popularity  float64 `json:"popularity"`
	Adult       bool    `json:"adult"`
}
