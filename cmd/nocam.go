package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func nocamWorker(p string) (FileBool, error) {
	x, err := internal.ParseFile(p)
	if err != nil {
		log.Fatalln(err)
		return FileBool{p, false}, err
	}
	if x.IFD0.Make == "" {
		return FileBool{p, true}, nil
	}
	if x.IFD0.Model == "" {
		return FileBool{p, true}, nil
	}
	return FileBool{p, false}, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "nocam",
		Short: "Find files without camera maker/model data",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			txt, err := internal.Dispatcher(
				imageFile,
				nocamWorker,
				listCollector,
				args[0],
				exclude,
			)
			cobra.CheckErr(err)
			fmt.Println(txt)
		},
	})
}
