package infrastructure

import (
	"github.com/go-resty/resty/v2"
)

const omdbAuthToken = "eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJhMGJkYWRhMmM5NTFhOTBiNmQxNjc4NjgyMTQ3NTRhMSIsInN1YiI6IjYxNWI5OTZjYzhhMmQ0MDAyYWMxMGM3YiIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.u9kwuL1lNbkvWKUhPqP6vVssioOMiv7a94Wa3cmOm4E"

func NewClient() *resty.Client {
	return resty.New().SetAuthToken(omdbAuthToken)
}
