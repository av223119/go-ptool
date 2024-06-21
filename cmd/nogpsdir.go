package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

var nogpsDirCmd = &cobra.Command{
	Use:   "nogpsdir",
	Short: "Files without GPS data, per directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txt, err := internal.Dispatcher(
			nogpsWorker,
			dirCollector,
			args[0],
			exclude,
		)
		cobra.CheckErr(err)
		fmt.Println(txt)
	},
}

func init() {
	rootCmd.AddCommand(nogpsDirCmd)
}
