package isatty_test

import (
	"os"
	"testing"

	"github.com/mattn/go-isatty"
)

func TestIsWriterTerminal(t *testing.T) {
	// test for non-panic
	t.Log("os.Stdout:", isatty.IsWriterTerminal(os.Stdout))
}

func TestIsWriterCygwinTerminal(t *testing.T) {
	// test for non-panic
	t.Log("os.Stdout:", isatty.IsWriterCygwinTerminal(os.Stdout))
}
