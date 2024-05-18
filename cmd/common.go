package cmd

import (
	"slices"
	"strings"
)

type FileBool struct {
	Filename string
	Result bool
}

func listCollector(input <-chan FileBool, output chan<- string) {
	defer close(output)
	res := []string{}
	for s := range input {
		if s.Result {
			res = append(res, s.Filename)
		}
	}
	slices.Sort(res)
	output <- strings.Join(res, "\n")
}
