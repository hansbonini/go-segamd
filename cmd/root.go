package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-segamd",
	Short: "A collection of tools to work with Sega Genesis / Mega Drive ROMs",
	Long: `A collection of tools to work with Sega Genesis / Mega Drive ROMs handling
graphics, audios, tilemaps, sprites, texts.`,
}

var RootCmd = rootCmd

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
