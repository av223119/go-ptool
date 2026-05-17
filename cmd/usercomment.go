package cmd

import (
	"fmt"
	"log"
	"unicode"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

func usercommentWorker(p string) (KVPair, error) {
	x, err := internal.ParseFile(p)
	if err != nil {
		log.Fatalln(err)
		return KVPair{}, err
	}
	if x.ExifIFD.UserComment != "" {
		rs := []rune(x.ExifIFD.UserComment)
		rs = rs[:min(40, len(rs))]
		for i, r := range rs {
			if unicode.IsSpace(r) {
				rs[i] = ' '
			}
		}
		return KVPair{p, string(rs)}, nil
	}
	return KVPair{}, nil
}


func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "usercomment",
		Short: "Files with UserComment",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			txt, err := internal.Dispatcher(
				imageFile,
				usercommentWorker,
				kvListCollector,
				args[0],
				exclude,
			)
			cobra.CheckErr(err)
			fmt.Println(txt)
		},
	})
}
