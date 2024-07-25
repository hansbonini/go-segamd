package cmd

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hansbonini/go-segamd/types"
	"github.com/hansbonini/go-segamd/types/generic"

	"github.com/spf13/cobra"
)

var pcmCmd = &cobra.Command{
	Use:   "pcm",
	Short: "Handle PCM Data on Sega Genesis / Mega Drive ROMs",
	Long:  `Handle PCM Data on Sega Genesis / Mega Drive ROMs`,
}

var pcm2wav = &cobra.Command{
	Use:        "pcm2wav",
	Short:      "Convert Sega Genesis / Mega Drive PCM Data to WAV",
	Long:       `Convert Sega Genesis / Mega Drive PCM Data to WAV`,
	Args:       cobra.MinimumNArgs(4),
	ArgAliases: []string{"input", "output", "channels", "samplerate"},
	ValidArgs:  []string{"input", "output", "channels", "samplerate"},
	Example:    `go-segamd pcm2wav input.pcm output.wav 1 4000`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.Fatal(err)
		}
		split := strings.Split(args[1], string(os.PathSeparator))
		path := strings.Join(split[:len(split)-1], string(os.PathSeparator))
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.MkdirAll(path, 0777); err != nil {
				log.Fatal(err)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var in *generic.ROM
		var out *os.File
		var channels int
		var samplerate int
		var err error
		if in, err = generic.NewROM(args[0]); err != nil {
			log.Fatal(err)
		}
		if out, err = os.Create(args[1]); err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		if channels, err = strconv.Atoi(args[2]); err != nil {
			log.Fatal(err)
		}
		if samplerate, err = strconv.Atoi(args[3]); err != nil {
			log.Fatal(err)
		}
		pcm := types.MDPCM{
			Channels:   channels,
			SampleRate: samplerate,
		}
		wav, err := pcm.ToWAV(in.Data)
		if err != nil {
			log.Fatal(err)
		}
		if _, err = out.Write(wav); err != nil {
			log.Fatal(err)
		}
	},
}

var wav2pcm = &cobra.Command{
	Use:     "wav2pcm",
	Short:   "Convert WAV Data to Sega Genesis / Mega Drive PCM Data",
	Long:    `Convert WAV Data to Sega Genesis / Mega Drive PCM Data`,
	Example: `go-segamd wav2pcm input.wav output.pcm 1 4000`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			log.Fatal(err)
		}
		split := strings.Split(args[1], string(os.PathSeparator))
		path := strings.Join(split[:len(split)-1], string(os.PathSeparator))
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.MkdirAll(path, 0777); err != nil {
				log.Fatal(err)
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var in *generic.ROM
		var out *os.File
		var channels int
		var samplerate int
		var err error
		if in, err = generic.NewROM(args[0]); err != nil {
			log.Fatal(err)
		}
		if out, err = os.Create(args[1]); err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		if channels, err = strconv.Atoi(args[2]); err != nil {
			log.Fatal(err)
		}
		if samplerate, err = strconv.Atoi(args[3]); err != nil {
			log.Fatal(err)
		}
		wav := types.MDPCM{
			Channels:   channels,
			SampleRate: samplerate,
		}
		pcm, err := wav.FromWAV(in.Data)
		if err != nil {
			log.Fatal(err)
		}
		if _, err = out.Write(pcm); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	pcmCmd.AddCommand(pcm2wav)
	pcmCmd.AddCommand(wav2pcm)
	rootCmd.AddCommand(pcmCmd)
}
