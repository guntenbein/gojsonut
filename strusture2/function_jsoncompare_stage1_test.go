package simple

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nsf/jsondiff"
)

func TestSortJsonCompareStage1(t *testing.T) {
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
			"Type": "circle",
			"Color": "black"
		  },
		  {
			"Type": "circle",
			"Color": "white"
		  },
		  {
			"Type": "square",
			"Color": "black"
		  },
		  {
			"Type": "square",
			"Color": "red"
		  },
		  {
			"Type": "square",
			"Color": "white"
		  }
		]
		`
	out := Sort(in)

	outJsonStr, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		t.Fatal("error marshaling package", err)
	}
	fmt.Println(string(outJsonStr))

	diffOpts := jsondiff.DefaultConsoleOptions()
	res, diff := jsondiff.Compare([]byte(expectedJsonStr), []byte(outJsonStr), &diffOpts)

	if res != jsondiff.FullMatch {
		t.Errorf("the expected result is not equal to what we have: %s", diff)
	}
}
