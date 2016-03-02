package entry

import "time"

type RepositoryDateStars struct {
	RepositoryName string `json:"-"`
	Date time.Time `json:"date"`
	Stars int16 `json:"stars"`
}
