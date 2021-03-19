package parser

import (
	"encoding/csv"
	"os"
	"strings"
	"time"

	"github.com/workmanager/calculator"
	"github.com/workmanager/entity"
)

var (
	FieldNumber = 14
)

// Reader reads the file with name *.csv
func Reader(filename string, fields int, start, end time.Time) ([]entity.InputData, map[string]entity.ProjectMetric, map[string]entity.TeamMemberMetric) {
	var list []entity.InputData
	projectMetric := make(map[string]entity.ProjectMetric)
	memberMetric := make(map[string]entity.TeamMemberMetric)

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
		startAct := entity.StringToTime(strings.Join([]string{gotData.StartDate, gotData.StartTime}, " "))
		endAct := entity.StringToTime(strings.Join([]string{gotData.EndDate, gotData.EndTime}, " "))
		if (startAct.Before(start) || startAct.After(end) || endAct.Before(start) || endAct.After(end)) && !(start.IsZero() || end.IsZero()) {
			continue
		}

		list = append(list, gotData)

		if _, ok := calculator.Projects[gotData.Project]; ok != true {
			calculator.Projects[gotData.Project] = entity.ProjectSalary{
				Salary:         1,
				OvertimeSalary: 1.5,
			}
		}
		calculator.UpdateMetricV2(gotData, projectMetric, memberMetric)

	}
	return list, projectMetric, memberMetric
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
