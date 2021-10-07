package common

type MovieResponse struct {
	Movie Movie `json:"movie"`
}

type MoviesResponse struct {
	Movies []Movie `json:"movies"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
