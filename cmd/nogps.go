package cmd

import (
	"fmt"

	"github.com/kolesa-team/goexiv"
	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal/collectors"
	"github.com/av223119/go-ptool/internal/dispatcher"
	"github.com/av223119/go-ptool/internal/image"
)

func nogps_worker(p string) (string, error) {
	exif, err := image.Exif(p)
	if err != nil {
		return "", err
	}

	// check GPS tags
	for _, field := range []string{"Exif.GPSInfo.GPSLatitude", "Exif.GPSInfo.GPSLongitude"} {
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

var nogpsCmd = &cobra.Command{
	Use:   "nogps",
	Short: "Find files without GPS data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txt, err := dispatcher.Run(
			nogps_worker,
			collectors.List_collector,
			args[0],
		)
		cobra.CheckErr(err)
		fmt.Println(txt)
	},
}

func init() {
	rootCmd.AddCommand(nogpsCmd)
}
