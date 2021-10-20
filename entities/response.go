package entities

// GetMovieResponse represents the response that will be returned if a movie was found
type GetMovieResponse struct {
	Movie Movie `json:"movie"`
}

// GetMoviesResponse represents the response that will be returned by the FindMovies method
type GetMoviesResponse struct {
	Movies []Movie `json:"movies"`
}

// ErrorResponse represents the response the app will return if an error occurred
type ErrorResponse struct {
	Message string `json:"message"`
}
