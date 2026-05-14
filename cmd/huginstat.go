package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func huginStatWorker(p string) (string, error) {
	x, err := internal.ParseFile(p)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	if strings.Contains(x.IFD0.Software, "Hugin") {
		return x.IFD0.Software, nil
	}
	return "", nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "huginstat",
		Short: "Hugin version statistics",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			txt, err := internal.Dispatcher(
				imageFile,
				huginStatWorker,
				countCollector,
				args[0],
				exclude,
			)
			cobra.CheckErr(err)
			fmt.Println(txt)
		},
	})
}
