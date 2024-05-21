package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tajtiattila/metadata/exif"
	"github.com/tajtiattila/metadata/exif/exiftag"

	"github.com/av223119/go-ptool/internal"
)

func nogpsWorker(p string) (FileBool, error) {
	x, err := internal.ParseFile(p)
	switch err {
	case exif.NotFound:
		return FileBool{p, true}, nil
	case nil:
		break
	default:
		return FileBool{p, false}, err
	}
	// get GPS
	// Can't use
	// if _, ok := x.GPSInfo(); !ok { }
	// because some exifs have SRational array of GPS nom/denom
	// Against the standard, but don't feel like patching the library
	for _, tagname := range []uint32{
		exiftag.GPSLatitude,
		exiftag.GPSLongitude,
		exiftag.GPSLatitudeRef,
		exiftag.GPSLongitudeRef,
	} {
		t := x.Tag(tagname)
		if !t.Valid() {
			return FileBool{p, true}, nil
		}
	}
	return FileBool{p, false}, nil
}

var nogpsCmd = &cobra.Command{
	Use:   "nogps",
	Short: "Find files without GPS data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txt, err := internal.Dispatcher(
			nogpsWorker,
			listCollector,
			args[0],
		)
		cobra.CheckErr(err)
		fmt.Println(txt)
	},
}

func init() {
	rootCmd.AddCommand(nogpsCmd)
}
