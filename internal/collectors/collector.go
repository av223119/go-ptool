package collectors

import (
	"slices"
	"strings"
)

func List_collector(input <-chan string, output chan<- string) {
	defer close(output)
	res := []string{}
	for s := range input {
		if s != "" {
			res = append(res, s)
		}
	}
	slices.Sort(res)
	output <- strings.Join(res, "\n")
}
