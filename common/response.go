package common

type MovieResponse struct {
	Movie Movie `json:"movie"`
}

type MoviesResponse struct {
	Movies []Movie `json:"movies"`
}

type errorResponse struct {
	Message string `json:"message"`
}
