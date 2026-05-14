package cmd

import (
	"fmt"
	"path"
	"sort"
	"strings"
)

type FileBool struct {
	Filename string
	Result   bool
}

type KVPair struct {
	Filename string
	Value    string
}

func rightmost(s string, n int) string {
	r := []rune(s)
	l := len(r)
	if len(r) <= n {
		return s
	}
	return string(r[max(0, l-n):])
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
	for i, k := range keys {
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

func kvListCollector(input <-chan KVPair, output chan<- string) {
	defer close(output)
	res := []KVPair{}
	for s := range input {
		if s.Filename != "" {
			res = append(res, s)
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i].Filename < res[j].Filename })
	strs := make([]string, len(res))
	for i, v := range res {
		strs[i] = fmt.Sprintf("%60v | %v", rightmost(v.Filename, 60), v.Value)
	}
	output <- strings.Join(strs, "\n")
}

func countCollector(input <-chan string, output chan<- string) {
	defer close(output)
	res := map[string]int{}
	for s := range input {
		if s == "" {
			continue
		}
		v, ok := res[s]
		if !ok {
			res[s] = 1
		} else {
			res[s] = v + 1
		}
	}
	keys := make([]string, 0, len(res))
	for k := range res {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return res[keys[i]] < res[keys[j]] })
	table := make([]string, len(res))
	for i, k := range keys {
		table[i] = fmt.Sprintf("%20v | %10v", k, res[k])
	}
	output <- strings.Join(table, "\n")
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

func imageFile(p string) bool {
	return strings.HasSuffix(p, ".jpg") || strings.HasSuffix(p, ".heic")
}

func anyFile(p string) bool {
	return true
}
