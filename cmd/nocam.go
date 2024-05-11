package cmd

import (
	"fmt"
	"os"

	"github.com/tajtiattila/metadata/exif"
	"github.com/tajtiattila/metadata/exif/exiftag"
	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func nocamWorker(p string) (string, error) {
	f, err := os.Open(p)
	if err != nil {
		return "", err
	}
	defer f.Close()
	// get exif
	x, err := exif.Decode(f)
	if err != nil {
		return "", nil
	}
	// check make and model
	for _, tagname := range []uint32{exiftag.Make, exiftag.Model} {
		t := x.Tag(tagname)
		if !t.Valid() {
			return p, nil
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
			nocamWorker,
			listCollector,
			args[0],
		)
		cobra.CheckErr(err)
		fmt.Println(txt)
	},
}

func init() {
	rootCmd.AddCommand(nocamCmd)
}
