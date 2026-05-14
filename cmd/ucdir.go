package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func ucboolWorker(p string) (FileBool, error) {
	x, err := internal.ParseFile(p)
	if err != nil {
		log.Fatalln(err)
		return FileBool{p, false}, err
	}
	if x.ExifIFD.UserComment != "" {
		return FileBool{p, true}, nil
	}
	return FileBool{p, false}, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "ucdir",
		Short: "Count files with UserComment, per directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			txt, err := internal.Dispatcher(
				imageFile,
				ucboolWorker,
				dirCollector,
				args[0],
				exclude,
			)
			cobra.CheckErr(err)
			fmt.Println(txt)
		},
	})
}
