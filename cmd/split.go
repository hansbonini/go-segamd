package cmd

import (
	"bufio"
	"fmt"
	"go-segamd/types/generic"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:        "split",
	Short:      "Split a range of bytes from a Sega Genesis / Mega Drive ROM",
	Long:       `Split a range of bytes from a Sega Genesis / Mega Drive ROM into a new file`,
	Args:       cobra.MinimumNArgs(2),
	ArgAliases: []string{"input", "split_list"},
	ValidArgs:  []string{"input", "split_list"},
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.Fatal(err)
		}
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var in *generic.ROM
		var split_list *os.File
		var err error
		if in, err = generic.NewROM(args[0]); err != nil {
			log.Fatal(err)
		}
		if split_list, err = os.Open(args[1]); err != nil {
			log.Fatal(err)
		}
		var lines []string
		scanner := bufio.NewScanner(split_list)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		var start_offset, end_offset int
		var filepath string
		for _, line := range lines {
			fmt.Sscanf(line, "0x%x,0x%x,%s", &start_offset, &end_offset, &filepath)
			filepath = strings.ReplaceAll(filepath, "\\", string(os.PathSeparator))
			filepath = strings.ReplaceAll(filepath, "/", string(os.PathSeparator))
			split_path := strings.Split(filepath, string(os.PathSeparator))
			fmt.Println(filepath)
			path := strings.Join(split_path[:len(split_path)-1], string(os.PathSeparator))
			if _, err := os.Stat(path); os.IsNotExist(err) {
				if err := os.MkdirAll(path, 0777); err != nil {
					log.Fatal(err)
				}
			}
			out, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()
			// read rom at start_offset
			out.Write(in.Data[start_offset:end_offset])
			out.Close()
		}
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)
}
