package cmd_test

import (
	"bytes"
	"testing"

	"github.com/hansbonini/go-segamd/cmd"
)

func TestSplitCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	c := cmd.RootCmd
	c.SetArgs([]string{"split"})
	c.SetOutput(buf)
	if err := c.Execute(); err != nil {
		if err.Error() != "requires at least 2 arg(s), only received 0" {
			t.Fatal(err)
		}
	}
}
