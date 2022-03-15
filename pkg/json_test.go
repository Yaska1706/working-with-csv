package pkg

import (
	"encoding/json"
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

	os.Remove(file)

}
