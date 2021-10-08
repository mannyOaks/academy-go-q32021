package infrastructure

type apiMovieResponse struct {
	ID               int     `json:"id"`
	Adult            bool    `json:"adult"`
	BackdropPath     string  `json:"backdrop_path"`
	Genres           []int   `json:"genre_ids"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Popularity       float64 `json:"popularity"`
	PosterPath       string  `json:"poster_path"`
	Title            string  `json:"title"`
	Video            bool    `json:"video"`
	VoteAvg          float32 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
}
