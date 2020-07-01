package simple

import (
	"testing"

	"github.com/guntenbein/gojsonut"
)

func TestSortJsonCompareStage3(t *testing.T) {
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
	expectedJsonStr := `
		[
		  {
			"Dimension": 2,
			"Type": "circle",
			"Color": "black"
		  },
		  {
			"Dimension": 2,
			"Type": "circle",
			"Color": "white"
		  },
		  {
			"Dimension": 2,
			"Type": "square",
			"Color": "black"
		  },
		  {
			"Dimension": 2,
			"Type": "square",
			"Color": "red"
		  },
		  {
			"Dimension": 2,
			"Type": "square",
			"Color": "white"
		  },
		  {
			"Dimension": 3,
			"Type": "cone",
			"Color": "black"
		  },
		  {
			"Dimension": 3,
			"Type": "cone",
			"Color": "white"
		  },
		  {
			"Dimension": 3,
			"Type": "cube",
			"Color": "black"
		  }
		]
		`
	out := Sort(in)

	gojsonut.JsonCompare(t, out, expectedJsonStr)
}
