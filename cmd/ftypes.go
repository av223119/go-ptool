package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func ftypesWorker(p string) (string, error) {
	s := strings.Split(p, ".")
	return s[len(s)-1], nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "ftypes",
		Short: "Overall filetype statistics",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			txt, err := internal.Dispatcher(
				anyFile,
				ftypesWorker,
				countCollector,
				args[0],
				exclude,
			)
			cobra.CheckErr(err)
			fmt.Println(txt)
		},
	})
}
