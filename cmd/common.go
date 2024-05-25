package cmd

import (
	"fmt"
	"path"
	"sort"
	"strings"
)

type FileBool struct {
	Filename string
	Result bool
}

type boolCounter map[bool]uint
type mapCounter map[string]boolCounter

func (m mapCounter) String() string {
	keys := []string{}
	for k, v := range m {
		if v[true] != 0 {
			keys = append(keys, k)
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]][true] < m[keys[j]][true]
	})
	res := make([]string, len(keys))
	for i, k := range(keys) {
		res[i] = fmt.Sprintf("%3v / %3v  %v", m[k][true], m[k][false], k)
	}
	return strings.Join(res, "\n")
}

func listCollector(input <-chan FileBool, output chan<- string) {
	defer close(output)
	res := []string{}
	for s := range input {
		if s.Result {
			res = append(res, s.Filename)
		}
	}
	sort.Strings(res)
	output <- strings.Join(res, "\n")
}

func dirCollector(input <-chan FileBool, output chan<- string) {
	defer close(output)
	res := mapCounter{}
	for s := range input {
		dirName := path.Dir(s.Filename)
		if res[dirName] == nil {
			res[dirName] = boolCounter{}
		}
		res[dirName][s.Result] += 1
	}
	output <- res.String()
}
