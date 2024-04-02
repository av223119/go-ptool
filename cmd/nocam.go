package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nocamCmd = &cobra.Command{
	Use:   "nocam",
	Short: "Find files without camera maker/model data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nocam called")
	},
}

func init() {
	rootCmd.AddCommand(nocamCmd)
}
