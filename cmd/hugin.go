package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func huginWorker(p string) (KVPair, error) {
	x, err := internal.ParseFile(p)
	if err != nil {
		log.Fatalln(err)
		return KVPair{}, err
	}
	if strings.Contains(x.IFD0.Software, "Hugin") {
		return KVPair{p, x.IFD0.Software}, nil
	}
	return KVPair{}, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "hugin",
		Short: "Files edited by Hugin",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			txt, err := internal.Dispatcher(
				imageFile,
				huginWorker,
				kvListCollector,
				args[0],
				exclude,
			)
			cobra.CheckErr(err)
			fmt.Println(txt)
		},
	})
}
