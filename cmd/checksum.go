package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/hansbonini/go-segamd/types/generic"

	"github.com/spf13/cobra"
)

// checksumCmd represents the checksum command
var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Generate a checksum of the ROM",
	Long:  `Generate a checksum of the ROM`,
}

var checkChecksumCmd = &cobra.Command{
	Use:        "check",
	Short:      "Check the checksum of the ROM",
	Long:       `Check the checksum of the ROM`,
	Args:       cobra.MinimumNArgs(3),
	ValidArgs:  []string{"input", "algorithm", "value"},
	ArgAliases: []string{"input", "algorithm", "value"},
	Example:    `go-segamd checksum check input.rom md5 6d7e6f1a6d7e6f1a6d7e6f1a`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.Fatal(err)
		}
		if !slices.Contains([]string{"md5", "sha1", "crc32"}, args[1]) {
			log.Fatal("Invalid checksum type. Valid algorithms: md5, sha1, crc32")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var in *generic.ROM
		var checksum string
		var err error
		if in, err = generic.NewROM(args[0]); err != nil {
			log.Fatal(err)
		}
		switch args[1] {
		case "md5":
			checksum = strings.ToUpper(fmt.Sprintf("%x", md5.Sum(in.Data)))
		case "sha1":
			checksum = strings.ToUpper(fmt.Sprintf("%x", sha1.Sum(in.Data)))
		case "crc32":
			checksum = strings.ToUpper(fmt.Sprintf("%x", crc32.ChecksumIEEE(in.Data)))
		}
		if strings.ToUpper(args[2]) != checksum {
			log.Fatal("Checksum does not match")
		}
		fmt.Println("Checksum matches")
	},
}

var getChecksumCmd = &cobra.Command{
	Use:        "get",
	Short:      "Get the checksum of the ROM",
	Long:       `Get the checksum of the ROM`,
	Args:       cobra.MinimumNArgs(2),
	ValidArgs:  []string{"input", "algorithm"},
	ArgAliases: []string{"input", "algorithm"},
	Example:    `go-segamd checksum get input.rom md5`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.Fatal(err)
		}
		if !slices.Contains([]string{"md5", "sha1", "crc32"}, args[1]) {
			log.Fatal("Invalid checksum type. Valid algorithms: md5, sha1, crc32")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var in *generic.ROM
		var checksum string
		var err error
		if in, err = generic.NewROM(args[0]); err != nil {
			log.Fatal(err)
		}
		switch args[1] {
		case "md5":
			checksum = strings.ToUpper(fmt.Sprintf("%x", md5.Sum(in.Data)))
		case "sha1":
			checksum = strings.ToUpper(fmt.Sprintf("%x", sha1.Sum(in.Data)))
		case "crc32":
			checksum = strings.ToUpper(fmt.Sprintf("%x", crc32.ChecksumIEEE(in.Data)))
		}
		fmt.Println(checksum)
	},
}

var listChecksumCmd = &cobra.Command{
	Use:        "list",
	Short:      "List all checksum of the ROM",
	Long:       `List all checksum of the ROM using all algorithms available`,
	Args:       cobra.MinimumNArgs(1),
	ValidArgs:  []string{"input"},
	ArgAliases: []string{"input"},
	Example:    `go-segamd checksum list input.rom`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.Fatal(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var in *generic.ROM
		var err error
		if in, err = generic.NewROM(args[0]); err != nil {
			log.Fatal(err)
		}
		fmt.Println("MD5:\t\t" + strings.ToUpper(fmt.Sprintf("%x", md5.Sum(in.Data))))
		fmt.Println("SHA1:\t\t" + strings.ToUpper(fmt.Sprintf("%x", sha1.Sum(in.Data))))
		fmt.Println("CRC32:\t\t" + strings.ToUpper(fmt.Sprintf("%x", crc32.ChecksumIEEE(in.Data))))
	},
}

func init() {
	checksumCmd.AddCommand(checkChecksumCmd)
	checksumCmd.AddCommand(getChecksumCmd)
	checksumCmd.AddCommand(listChecksumCmd)
	rootCmd.AddCommand(checksumCmd)
}
