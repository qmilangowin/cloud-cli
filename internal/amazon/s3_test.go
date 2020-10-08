package amazon_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	c, err = root.ExecuteC()
	return c, buf.String(), err
}

func TestS3all(t *testing.T) {
	var rootCmdArgs []string
	s3list := &cobra.Command{
		Use:  "s3",
		Args: cobra.ExactArgs(1),
		Run:  func(_ *cobra.Command, args []string) { rootCmdArgs = args },
	}
	output, err := executeCommand(s3list, "all")

	if output != "" {
		t.Errorf("Unexpected output: %v ", output)
	}

	if err != nil {
		t.Errorf("Unexpected error: %v ", err)
	}

	got := strings.Join(rootCmdArgs, " ")
	expected := "all"
	if got != expected {
		t.Errorf("app s3 all expected: %q, got: %q", got, expected)
	}

}
