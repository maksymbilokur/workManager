package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"taskmanager/entity"
)

// Reader reads the file with name *.csv
func Reader(filename string, fields int) []entity.InputData {
	var list []entity.InputData
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = fields
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break

		}
		list = append(list, *getStruct(record))
	}
	return list
}

func getStruct(elem []string) *entity.InputData {
	return &entity.InputData{
		User:        elem[entity.User],
		Email:       elem[entity.Email],
		Client:      elem[entity.Client],
		Project:     elem[entity.Project],
		Task:        elem[entity.Task],
		Description: elem[entity.Description],
		Billable:    elem[entity.Billable],
		StartDate:   elem[entity.StartDate],
		StartTime:   elem[entity.StartTime],
		EndDate:     elem[entity.EndDate],
		EndTime:     elem[entity.EndTime],
		Duration:    elem[entity.Duration],
		Tag:         elem[entity.Tag],
	}
}
