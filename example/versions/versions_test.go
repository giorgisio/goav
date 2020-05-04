package main

import (
	"testing"
)

func TestVersions(t *testing.T) {
	// Check this doesn't cause SEGV or link error
	PrintVersions()
}
