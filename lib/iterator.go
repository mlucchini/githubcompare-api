package lib

import (
	"io"
	"bufio"
)

func GroupLinesIterator(reader io.Reader, nbLinesPerGroup int) (<- chan []string) {
	channel := make(chan []string)
	go func() {
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
	}()
	return channel
}