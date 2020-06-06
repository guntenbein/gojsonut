package simple

import (
	"encoding/json"
	"testing"

	"github.com/nsf/jsondiff"
)

func TestSortJsonCompareStage2(t *testing.T) {
	in := []Figure{
		{"circle", "white"},
		{"square", "black"},
		{"circle", "black"},
		{"square", "white"},
		{"square", "red"},
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
	//fmt.Println(string(outJsonStr))

	diffOpts := jsondiff.DefaultConsoleOptions()
	res, diff := jsondiff.Compare([]byte(expectedJsonStr), []byte(outJsonStr), &diffOpts)

	if res != jsondiff.FullMatch {
		t.Errorf("the expected result is not equal to what we have: %s", diff)
	}
}
