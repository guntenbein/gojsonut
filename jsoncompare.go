package gojsonut

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nsf/jsondiff"
)

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

	if res != jsondiff.ßFullMatch {
		t.Errorf("the expected result is not equal to what we have: \n %s", diff)
	}
}
