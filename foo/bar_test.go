package foo_test

import (
	"testing"

	"github.com/shachartal/intro-to-go/foo"
)

func TestAdd(t *testing.T) {
	if foo.Add(1, 2) != 3 {
		t.Error("adding doesn't work properly")
	}
	if foo.Add(2, 2) != 4 {
		t.Error("adding doesn't work properly")
	}
}
