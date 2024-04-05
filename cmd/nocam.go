package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/kolesa-team/goexiv"
	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal/dispatcher"
	"github.com/av223119/go-ptool/internal/image"
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
	exif, err := image.Exif(p)
	if err != nil {
		return "", err
	}

	// check make and model
	for _, field := range []string{"Exif.Image.Make", "Exif.Image.Model"} {
		_, err := exif.GetString(field)
		if err != nil {
			if err == goexiv.ErrMetadataKeyNotFound {
				return p, nil
			}
			return "", err
		}
	}
	return "", nil
}

var nocamCmd = &cobra.Command{
	Use:   "nocam",
	Short: "Find files without camera maker/model data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txt, err := dispatcher.Run(
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
