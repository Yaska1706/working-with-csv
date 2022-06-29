package pkg

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"

	"golang.org/x/text/encoding/unicode"
)

type employee struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func GetDataFromCSV(filename string) [][]string {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("CSV file opened")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(unicode.UTF8.NewDecoder().Reader(csvFile)).ReadAll()
	if err != nil {
		fmt.Println("failed to read csv", err)
		return nil
	}
	return csvLines
}

func StoreCSVData(csvlines [][]string) []employee {
	employees := make([]employee, 0)

	for i, line := range csvlines {
		if i > 0 {
			var emp employee
			for j, field := range line {
				if j == 0 {
					emp.Name = field
				} else if j == 1 {
					emp.Age = field
				} else if j == 2 {
					emp.City = field
				}
			}
			employees = append(employees, emp)
		}

	}
	return employees

}

func WriteToCSV(filename string, data [][]string) {
	csvFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	for _, row := range data {
		if err := csvWriter.Write(row); err != nil {
			fmt.Println("something happened:", err)
		}
	}
	csvWriter.Flush()

}

func WriteAlltoCSV(filename string, data [][]string) {
	csvFile, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)

	}

	csvFile.Close()

	cW := csv.NewWriter(csvFile)
	if err := cW.WriteAll(data); err != nil {
		log.Fatal(err)
	}
}

func CheckExist(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}

func DeleteFile(filename string) {
	if err := os.Remove(filename); err != nil {
		fmt.Println("failed to delete file:", err)
	}
}
