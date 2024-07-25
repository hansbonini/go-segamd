package cmd_test

import (
	"bytes"
	"testing"

	"github.com/hansbonini/go-segamd/cmd"
)

func TestChecksumCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	c := cmd.RootCmd
	c.SetArgs([]string{"checksum"})
	c.SetOutput(buf)
	if err := c.Execute(); err != nil {
		t.Fatal(err)
	}
}
