package model

import (
	"fmt"
	"strings"
	"io"
	"encoding/json"
	"strconv"
)

type RepositoryStats struct {
	RepositoryName string `json:"repository"`
	Stars []int `json:"stars" datastore:",noindex"`
}

func (self *RepositoryStats) String() string {
	return fmt.Sprintf("%s,%+v", self.RepositoryName, self.Stars)
}

func (self *RepositoryStats) Json(w io.Writer)  {
	encoder := json.NewEncoder(w)
	encoder.Encode(self)
}

func (self *RepositoryStats) Parse(csv string) error {
	// Format: org/repo,1;2;3;4;5;6;7;8;9;10;11;12

	values := strings.Split(csv, ",")

	self.RepositoryName = values[0]

	stars := strings.Split(values[1], ";")
	self.Stars = make([]int, 0, len(stars))
	for _, s := range stars {
		i, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		self.Stars = append(self.Stars, i)
	}
	return nil
}