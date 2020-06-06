package simple

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nsf/jsondiff"
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

	JsonCompare(t, out, expectedJsonStr, false)
}

func JsonCompare(t *testing.T, out interface{}, expectedJsonStr string, print bool) {
	outJsonStr, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		t.Fatal("error marshaling package", err)
	}
	if print {
		fmt.Println(string(outJsonStr))
	}
	diffOpts := jsondiff.DefaultConsoleOptions()
	res, diff := jsondiff.Compare([]byte(expectedJsonStr), []byte(outJsonStr), &diffOpts)

	if res != jsondiff.FullMatch {
		t.Errorf("the expected result is not equal to what we have: \n %s", diff)
	}
}
