package simple

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	in := []Figure{
		{"circle", "white"},
		{"square", "black"},
		{"circle", "black"},
		{"square", "white"},
		{"square", "red"},
	}
	expected := []Figure{
		{"circle", "black"},
		{"circle", "white"},
		{"square", "black"},
		{"square", "red"},
		{"square", "white"},
	}
	out := Sort(in)
	if !reflect.DeepEqual(expected, out) {
		t.Fatalf("the expected result %v is not equal to what we have %v", expected, out)
	}
}
