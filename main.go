package main

import (
	"encoding/json"
	"fmt"

	"github.com/yaska1706/working-with-csv/pkg"
)

func main() {
	filename := "employee1.csv"

	data := [][]string{
		{"name", "age", "city"},
		{"Smith", "23", "Nairobi"},
		{"William", "34", "Nakuru"},
		{"Rose", "41", "Mombasa"},
	}
	b, err := pkg.CheckExist(filename)
	if err != nil {
		fmt.Println("failed check", err)
	}
	if b {
		csvdata := pkg.GetDataFromCSV(filename)
		employeesdata := pkg.StoreCSVData(csvdata)
		jsonData, err := json.Marshal(employeesdata)
		if err != nil {
			fmt.Println("somithing happened,", err)

		}
		fmt.Println(string(jsonData))
	} else {

		pkg.WriteToCSV(filename, data)
	}
}
