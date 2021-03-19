package parser

import (
	"encoding/csv"
	"os"

	"github.com/workmanager/calculator"
	"github.com/workmanager/entity"
)

func ReaderWithoutMetric(filename string, fields int) []entity.InputData {
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
			break
		}

		gotData := *getStruct(record)
		list = append(list, gotData)

		if _, ok := calculator.Projects[gotData.Project]; ok != true {
			calculator.Projects[gotData.Project] = entity.ProjectSalary{
				Salary:         1,
				OvertimeSalary: 1.5,
			}
		}

	}
	return list
}
