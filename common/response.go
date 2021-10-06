package common

type MovieResponse struct {
	Message string `json:"message"`
	Movie   Movie  `json:"movie"`
}

type MoviesResponse struct {
	Message string  `json:"message"`
	Movies  []Movie `json:"movies"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
