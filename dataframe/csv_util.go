package dataframe

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"web_scraping/exceptions"
)

func ToCSV(dataList []interface{}, path string) {
	if dataList == nil || len(dataList) == 0 {
		return
	}

	// find headers
	var headers []string
	for i := 0; i < reflect.TypeOf(dataList[0]).NumField(); i++ {
		headers = append(headers, reflect.TypeOf(dataList[0]).Field(i).Name)
	}

	// handle rows
	var rows [][]string
	for _, data := range dataList {
		if reflect.ValueOf(data).Kind() == reflect.Struct {
			var row []string
			for i := 0; i < reflect.TypeOf(data).NumField(); i++ {
				v := reflect.ValueOf(data).Field(i)
				if v.Kind() != reflect.String {
					row = append(row, fmt.Sprintf("%v", v))
				} else {
					row = append(row, reflect.ValueOf(data).Field(i).String())
				}
			}
			rows = append(rows, row)
		}
	}

	// write to csv file
	file, err := os.Create(path)
	exceptions.HandleError(err)
	defer func(file *os.File) {
		err := file.Close()
		exceptions.HandleError(err)
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(headers)
	exceptions.HandleError(err)

	for _, row := range rows {
		err = writer.Write(row)
		exceptions.HandleError(err)
	}
}

func ReadCSV(path string) []map[string]string {
	// read from csv
	file, err := os.Open(path)
	exceptions.HandleError(err)

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	exceptions.HandleError(err)

	if len(rows) <= 1 {
		panic("insufficient rows in csv file")
	}

	headers := rows[0]
	data := rows[1:]
	var result []map[string]string

	for _, row := range data {
		m := make(map[string]string)
		for i, cell := range row {
			m[headers[i]] = cell
		}
		result = append(result, m)
	}
	return result
}
