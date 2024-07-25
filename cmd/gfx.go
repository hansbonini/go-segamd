package cmd

import (
	"go-segamd/types"
	"go-segamd/types/generic"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var gfxCmd = &cobra.Command{
	Use:   "gfx",
	Short: "Handle graphics on Sega Genesis / Mega Drive ROMs",
	Long:  `Handle graphics on Sega Genesis / Mega Drive ROMs`,
}

var gfx2pngCmd = &cobra.Command{
	Use:        "gfx2png",
	Short:      "Convert Sega Genesis / Mega Drive graphics to PNG",
	Long:       `Convert Sega Genesis / Mega Drive graphics to PNG`,
	Args:       cobra.MinimumNArgs(3),
	ValidArgs:  []string{"input", "output", "palette"},
	ArgAliases: []string{"input", "output", "palette"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.Fatal(err)
		}
		split := strings.Split(args[1], string(os.PathSeparator))
		if len(split[:len(split)-1]) > 0 {
			path := strings.Join(split[:len(split)-1], string(os.PathSeparator))
			if _, err := os.Stat(path); os.IsNotExist(err) {
				if err := os.MkdirAll(path, 0777); err != nil {
					log.Fatal(err)
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var in, pal *generic.ROM
		var out *os.File
		var err error
		if in, err = generic.NewROM(args[0]); err != nil {
			log.Fatal(err)
		}
		if out, err = os.Create(args[1]); err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		if pal, err = generic.NewROM(args[2]); err != nil {
			log.Fatal(err)
		}

		tiles := types.NewMDTiles(in.Data, 16, 4)
		palette := types.NewMDPalette(pal.Data)
		err = png.Encode(out, tiles.ToPNG(*palette))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	gfxCmd.AddCommand(gfx2pngCmd)
	rootCmd.AddCommand(gfxCmd)
}
