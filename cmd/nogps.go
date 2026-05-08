package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func nogpsWorker(p string) (FileBool, error) {
	x, err := internal.ParseFile(p)
	if err != nil {
		log.Fatalln(err)
		return FileBool{p, false}, err
	}
	if x.GPS.Latitude() == 0 && x.GPS.Longitude() == 0 {
		return FileBool{p, true}, nil
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
			exclude,
		)
		cobra.CheckErr(err)
		fmt.Println(txt)
	},
}

func init() {
	rootCmd.AddCommand(nogpsCmd)
}
