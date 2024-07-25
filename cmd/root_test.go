package cmd_test

import (
	"bytes"
	"testing"

	"github.com/hansbonini/go-segamd/cmd"
)

func TestRootCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	c := cmd.RootCmd
	c.SetOutput(buf)
	if err := c.Execute(); err != nil {
		t.Fatal(err)
	}

	c.Flags().Set("toggle", "true")
	if err := c.Execute(); err != nil {
		t.Fatal(err)
	}
}
