package main

import "testing"

func TestCommit(t *testing.T) {
	if commit == "" {
		t.Error("expected a value for commit, got empty string")
	}
}
