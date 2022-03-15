package pkg

import (
	"fmt"
	"os"
	"testing"
)

func createtemp(filename string) string {
	f, err := os.CreateTemp("", filename)
	if err != nil {
		fmt.Print(err)
	}
	return f.Name()
}

func TestDeleteFile(t *testing.T) {
	file := createtemp("test")

	DeleteFile(file)
	b, err := CheckExist(file)
	if err != nil {
		t.Error(err)
	}

	if b {
		t.Errorf("file was not deleted")
	}

}
