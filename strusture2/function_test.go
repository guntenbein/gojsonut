package simple

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	in := []Figure{
		{2, "circle", "white"},
		{2, "square", "black"},
		{3, "cone", "black"},
		{2, "circle", "black"},
		{2, "square", "white"},
		{3, "cone", "white"},
		{2, "square", "red"},
		{3, "cube", "black"},
	}
	expected := []Figure{
		{2, "circle", "black"},
		{2, "circle", "white"},
		{2, "square", "black"},
		{2, "square", "red"},
		{2, "square", "white"},
		{3, "cone", "black"},
		{3, "cone", "white"},
		{3, "cube", "black"},
	}
	out := Sort(in)
	if !reflect.DeepEqual(expected, out) {
		t.Fatalf("the expected result %v is not equal to what we have %v", expected, out)
	}
}
