package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestSaveToJsonFile(t *testing.T) {
	file := createtemp("test.json")

	test := map[string]string{
		"name": "new",
		"type": "greet",
	}
	data, err := json.Marshal(test)
	if err != nil {
		t.Error(err)
	}

	SaveToJsonFile(string(data))
	tempdata, err := os.ReadFile("employee.json")
	if err != nil {
		fmt.Print(err)
	}
	if string(tempdata) != string(data) {
		t.Errorf("not equal , wanted %q got %q", string(tempdata), string(data))
	}

	if err := os.Remove("employee.json"); err != nil {
		fmt.Print(err)
	}
	os.Remove(file)

}
