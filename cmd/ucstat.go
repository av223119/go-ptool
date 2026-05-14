package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func ucstatWorker(p string) (string, error) {
	x, err := internal.ParseFile(p)
	if err != nil {
		log.Fatalln(err)
		return "parse_error", err
	}
	if x.ExifIFD.UserComment != "" {
		return "present", nil
	}
	return "absent", nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "ucstat",
		Short: "Overall UserComment statistics",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			txt, err := internal.Dispatcher(
				imageFile,
				ucstatWorker,
				countCollector,
				args[0],
				exclude,
			)
			cobra.CheckErr(err)
			fmt.Println(txt)
		},
	})
}
