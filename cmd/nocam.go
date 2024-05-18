package cmd

import (
	"fmt"

	"github.com/tajtiattila/metadata/exif"
	"github.com/tajtiattila/metadata/exif/exiftag"
	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func nocamWorker(p string) (FileBool, error) {
	x, err := internal.ParseFile(p)
	switch err {
	case exif.NotFound:
		return FileBool{p, true}, nil
	case nil:
		break
	default:
		return FileBool{p, false}, err
	}
	// check make and model
	for _, tagname := range []uint32{exiftag.Make, exiftag.Model} {
		t := x.Tag(tagname)
		if !t.Valid() {
			return FileBool{p, true}, nil
		}
	}
	return FileBool{p, false}, nil
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
