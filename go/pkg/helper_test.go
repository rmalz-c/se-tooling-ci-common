package helper

import "testing"

func TestHelper(t *testing.T) {
	if !Help() {
		t.Error("Exected Hlper() to just work")
	}
}
