package project

import (
	"testing"

)

func TestNewRepo(t *testing.T){
	NewProject("driving-order")
}

func TestSync(t *testing.T) {
	// repo := NewProject("driving-order")
	// repo.Sync()
}

func TestDiff(t *testing.T) {
	// repo := NewProject("driving-order")
	// repo.Diff("a54e3d62a207580061da185a6f1ee267f6d22f94","a351e177f66c3bb20efbbadc03a0788cd92ab4ff")
}