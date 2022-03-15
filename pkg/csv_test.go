package pkg

import (
	"fmt"
	"os"
	"testing"
)

func createtemp() string {
	f, err := os.CreateTemp("", "test")
	if err != nil {
		fmt.Print(err)
	}
	return f.Name()
}

func TestDeleteFile(t *testing.T) {
	file := createtemp()

	DeleteFile(file)
	b, err := checkExist(file)
	if err != nil {
		t.Error(err)
	}

	if b {
		t.Errorf("file was not deleted")
	}

}
