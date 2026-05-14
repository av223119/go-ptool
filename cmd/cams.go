package cmd

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/av223119/go-ptool/internal"
)

type MakerModel struct {
	Maker string
	Model string
}

func camsWorker(p string) (MakerModel, error) {
	x, err := internal.ParseFile(p)
	if err != nil {
		log.Fatalln(err)
		return MakerModel{"-", "-"}, err
	}
	return MakerModel{x.IFD0.Make, x.IFD0.Model}, nil
}

func camsCollector(input <-chan MakerModel, output chan<- string) {
	defer close(output)
	res := map[string]map[string]uint{}
	for mm := range input {
		_, ok := res[mm.Maker]
		if !ok {
			res[mm.Maker] = map[string]uint{}
		}
		res[mm.Maker][mm.Model]++
	}
	var out strings.Builder
	mas := []string{}
	for m := range res {
		mas = append(mas, m)
	}
	sort.Strings(mas)
	for _, ma := range mas {
		mos := []string{}
		for m := range res[ma] {
			mos = append(mos, m)
		}
		sort.Strings(mos)
		for _, mo := range mos {
			fmt.Fprintf(&out, "%25s | %40s | %d\n", ma, mo, res[ma][mo])
		}
	}
	output <- out.String()
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "cams",
		Short: "Print camera maker and model stats",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			txt, err := internal.Dispatcher(
				imageFile,
				camsWorker,
				camsCollector,
				args[0],
				exclude,
			)
			cobra.CheckErr(err)
			fmt.Println(txt)
		},
	})
}
