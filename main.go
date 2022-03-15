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
	errCheck(err)
	if b {
		jsonResult := convertCSVtoJSON(filename)
		fmt.Print(jsonResult)
	} else {

		pkg.WriteToCSV(filename, data)
	}
}

func convertCSVtoJSON(filename string) string {

	csvdata := pkg.GetDataFromCSV(filename)
	employeesdata := pkg.StoreCSVData(csvdata)
	jsonData, err := json.Marshal(employeesdata)
	errCheck(err)

	return string(jsonData)

}

func errCheck(err error) {
	if err != nil {
		fmt.Println("somithing happened,", err)

	}
}
