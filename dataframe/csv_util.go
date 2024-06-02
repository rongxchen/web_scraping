package dataframe

import (
	"encoding/csv"
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
				row = append(row, reflect.ValueOf(data).Field(i).String())
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

func ReadCSV() {

}
