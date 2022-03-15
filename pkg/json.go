package pkg

import (
	"fmt"
	"os"
)

func SaveToJsonFile(data string) {
	f, err := os.Create("employee.json")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	f.Write([]byte(data))
	f.Sync()

}
