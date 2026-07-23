package isatty

import "testing"

func TestIsTerminalFileNil(t *testing.T) {
	if IsTerminalFile(nil) {
		t.Fatal("nil should not be terminal")
	}
	if IsCygwinTerminalFile(nil) {
		t.Fatal("nil should not be cygwin terminal")
	}
}
