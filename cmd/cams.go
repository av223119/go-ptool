package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tajtiattila/metadata/exif"
	"github.com/tajtiattila/metadata/exif/exiftag"

	"github.com/av223119/go-ptool/internal"
)

type MakerModel struct {
	Maker string
	Model string
}

func mytrim (s string) string {
	return strings.Trim(strings.TrimSpace(s), "\x00")
}

func camsWorker(p string) (MakerModel, error) {
	x, err := internal.ParseFile(p)
	switch err {
	case exif.NotFound:
		return MakerModel{"-", "-"}, nil
	case nil:
		break
	default:
		return MakerModel{"-", "-"}, err
	}
	ma, mo := "-", "-"
	t := x.Tag(exiftag.Make)
	if t.Valid() {
		ma, _ = t.Ascii()
	}
	t = x.Tag(exiftag.Model)
	if t.Valid() {
		mo, _ = t.Ascii()
	}
	return MakerModel{mytrim(ma), mytrim(mo)}, nil
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
		for _, mo := range mos {
			fmt.Fprintf(&out, "%25s | %40s | %d\n", ma, mo, res[ma][mo])
		}
	}
	output <- out.String()
}

var camsCmd = &cobra.Command{
	Use:   "cams",
	Short: "Print camera maker and model stats",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		txt, err := internal.Dispatcher(
			camsWorker,
			camsCollector,
			args[0],
			exclude,
		)
		cobra.CheckErr(err)
		fmt.Println(txt)
	},
}

func init() {
	rootCmd.AddCommand(camsCmd)
}
