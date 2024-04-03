package cmd

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func collector(input <-chan string, output chan<- string) {
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

func worker(p string) (string, error) {
	f, err := os.Open(p)
	if err != nil {
		return "", err
	}
	defer f.Close()
	// get exif
	x, err := exif.Decode(f)
	if err != nil {
		// no exif at all
		if err == io.EOF {
			return p, nil
		}
		// report critical errors
		if exif.IsCriticalError(err) {
			return "", err
		}
		// ignore all the rest
	}
	// check make and model
	for _, field := range []exif.FieldName{exif.Make, exif.Model} {
		_, err = x.Get(field)
		if err != nil {
			if exif.IsTagNotPresentError(err) {
				return p, nil
			}
			// ignore all other errors
		}
	}
	return "", nil
}

var nocamCmd = &cobra.Command{
	Use:   "nocam",
	Short: "Find files without camera maker/model data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txt, err := internal.Dispatcher(
			worker,
			collector,
			args[0],
		)
		cobra.CheckErr(err)
		fmt.Println(txt)
	},
}

func init() {
	rootCmd.AddCommand(nocamCmd)
}
