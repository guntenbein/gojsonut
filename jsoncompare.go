package gojsonut

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/nsf/jsondiff"
)

func JsonCompare(t *testing.T, result interface{}, expectedJsonStr string, print bool) {
	outJsonStr, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		t.Fatal("error marshaling the result: ", err)
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
