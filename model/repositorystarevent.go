package model

import (
	"time"
	"fmt"
	"strings"
	"strconv"
)

type RepositoryStarEvent struct {
	RepositoryName string `json:"-"`
	Date time.Time `json:"date"`
	Stars int `json:"stars"`
}

const YearMonthDayFormat = "2006-01-02"

func (self *RepositoryStarEvent) String() string {
	return fmt.Sprintf("%s,%s,%d", self.RepositoryName, self.Date.Format(YearMonthDayFormat), self.Stars)
}

func (self *RepositoryStarEvent) Parse(csv string) (error) {
	values := strings.Split(csv, ",")

	self.RepositoryName = values[0]

	if date, err := time.Parse(YearMonthDayFormat, values[1]); err != nil {
		return err
	} else {
		self.Date = date
	}

	if stars, err := strconv.Atoi(values[2]); err != nil {
		return err
	} else {
		self.Stars = stars
	}

	return nil
}