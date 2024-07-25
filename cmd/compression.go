/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-segamd/types"
	"go-segamd/types/generic"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// compressionCmd represents the compression command
var compressionCmd = &cobra.Command{
	Use:   "compression",
	Short: "Handle Sega Genesis / Mega Drive ROMs compression",
	Long:  `Handle Sega Genesis / Mega Drive ROMs compression`,
}

var segardCompressionCmd = &cobra.Command{
	Use:   "segard",
	Short: "Handle Sega Genesis / Mega Drive ROMs \"SEGARD\" compression",
	Long: `Handle Sega Genesis / Mega Drive ROMs \"SEGARD\" compression
Games where this compression is found:
	- [SMD] Alex Kidd in Enchanted Castle
	- [SMD] Altered Beast
	- [SMD] Columns
	- [SMD] Golden Axe
	- [SMD] Hokuto no Ken: Shin Seikimatsu Kyuuseishu Densetsu
	- [SMD] Last Battle
	- [SMD] Osomatsu-kun - Hachamecha Gekijou
	- [SMD] World Championship Soccer`,
	Args:       cobra.MinimumNArgs(3),
	ValidArgs:  []string{"mode", "input", "output"},
	ArgAliases: []string{"mode", "input", "output"},
	PreRun: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "decompress":
		case "compress":
		default:
			log.Fatal("Invalid mode. Valid modes: decompress, compress")
		}

		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			log.Fatal(err)
		}

		split := strings.Split(args[2], string(os.PathSeparator))
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
		var in *generic.ROM
		var out *os.File
		var err error
		var data []byte
		if in, err = generic.NewROM(args[1]); err != nil {
			log.Fatal(err)
		}
		if out, err = os.Create(args[2]); err != nil {
			log.Fatal(err)
		}
		compressor := types.NewMDCompressor("SEGARD", *in)
		switch args[0] {
		case "decompress":
			data = compressor.Unmarshal()
		case "compress":
			data = compressor.Marshal()
		}
		if len(data) > 0 {
			if _, err = out.Write(data); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatalf("Unable to %s data", args[0])
		}
	},
}

func init() {
	compressionCmd.AddCommand(segardCompressionCmd)
	rootCmd.AddCommand(compressionCmd)
}
