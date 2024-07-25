package cmd_test

import (
	"bytes"
	"testing"

	"github.com/hansbonini/go-segamd/cmd"
)

func TestPcmCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	c := cmd.RootCmd
	c.SetArgs([]string{"pcm"})
	c.SetOutput(buf)
	if err := c.Execute(); err != nil {
		t.Fatal(err)
	}
}