package internal

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"runtime"
)

func worker[Result any](
	filenameChan <-chan string,
	outputChan chan<- Result,
	doneChan chan<- bool,
	f func(string) Result,
) {
	// As long as filenameChan is open, read filename, call f,
	// push result to the outputChan
	for path := range filenameChan {
		outputChan <- f(path)
	}
	doneChan <- true
}

func Dispatcher[Result any](
	workerFunc func(string) Result,
	collectorFunc func(<-chan Result, chan<- string),
	rootDir string,
) (string, error) {
	filenameChan := make(chan string)
	resultChan := make(chan Result)
	doneChan := make(chan bool)
	textOut := make(chan string)
	nWorkers := runtime.NumCPU()

	// spawn N workers and one collector
	for i := 0; i < nWorkers; i++ {
		go worker(filenameChan, resultChan, doneChan, workerFunc)
	}
	go collectorFunc(resultChan, textOut)

	// walk the tree rooted at rootDir
	err := filepath.WalkDir(rootDir,
		func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return fmt.Errorf("Dispatcher: %w", err)
			}
			if !d.IsDir() {
				filenameChan <- p
			}
			return nil
		})

	// Close filenameChan, wait for all workers to finish
	close(filenameChan)
	for i := 0; i < nWorkers; i++ {
		<-doneChan
	}

	// Close resultChan, indicating that collector should return text
	close(resultChan)
	return <-textOut, err
}
