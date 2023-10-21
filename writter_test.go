package isatty_test

import (
	"os"
	"testing"

	"github.com/mattn/go-isatty"
)

func TestIsWriterTerminal(t *testing.T) {
	// test for non-panic
	isatty.IsWriterTerminal(os.Stdout)
}

func TestIsWriterCygwinTerminal(t *testing.T) {
	// test for non-panic
	isatty.IsWriterCygwinTerminal(os.Stdout)
}
