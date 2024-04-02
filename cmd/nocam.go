package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/av223119/go-ptool/internal"

	"github.com/spf13/cobra"
)

func collector(input <-chan string, output chan<- string) {
	defer close(output)
	res := []string{}
	for s := range input {
		res = append(res, s)
	}
	slices.Sort(res)
	output <- strings.Join(res, "\n")
}

func worker(p string) string {
	return p
}

var nocamCmd = &cobra.Command{
	Use:   "nocam",
	Short: "Find files without camera maker/model data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txt, err := internal.Dispatcher(
			worker,
			collector,
			args[0],
		)
		cobra.CheckErr(err)
		fmt.Println(txt)
	},
}

func init() {
	rootCmd.AddCommand(nocamCmd)
}
