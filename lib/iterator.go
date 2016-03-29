package lib

import (
	"io"
	"bufio"
)

func groupLinesIterator(reader io.Reader, nbLinesPerGroup int) (<- chan []string) {
	channel := make(chan []string)
	go split(reader, nbLinesPerGroup, channel)
	return channel
}

func split(reader io.Reader, nbLinesPerGroup int, channel chan []string) {
	scanner := bufio.NewScanner(reader)
	group := make([]string, 0, nbLinesPerGroup)
	for scanner.Scan() {
		group = append(group, scanner.Text())
		if (len(group) == nbLinesPerGroup) {
			channel <- group
			group = make([]string, 0, nbLinesPerGroup)
		}
	}
	if len(group) > 0 {
		channel <- group
	}
	close(channel)
}