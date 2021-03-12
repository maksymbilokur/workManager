package entity

import (
	"fmt"
	"strings"
	"time"
)

const (
	User = iota
	Email
	Client
	Project
	Task
	Description
	Billable
	StartDate
	StartTime
	EndDate
	EndTime
	Duration
	Tag
)

type InputData struct {
	User        string
	Email       string
	Client      string
	Project     string
	Task        string
	Description string
	Billable    string
	StartDate   string
	StartTime   string
	EndDate     string
	EndTime     string
	Duration    string
	Tag         string
}
type ActivityData struct {
	User     string
	Project  string
	Task     string
	Billable string
	Start    time.Time
	End      time.Time
	Duration float32
}

type ProjectSalary struct {
	Salary         float32
	OvertimeSalary float32
}

type ProjectMetric struct {
	WorkingHours float32
	TeamMembers  []string
}

type TeamMemberMetric struct {
	Salary            float32
	TotalWorkingHours float32
}

func StringToData(str string) time.Time {
	layout := "1/_2/2006 15:04:05"
	if t, err := time.Parse(layout, str); err == nil {
		return t
	}
	return time.Time{}
}

func StringToDuration(d string) float32 {
	sep := strings.Split(d, ":")
	hour := sep[0] + "h" + sep[1] + "m" + sep[2] + "s"
	res, err := time.ParseDuration(hour)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return float32(res.Hours())
}
func InputToActivity(data InputData) ActivityData {
	return ActivityData{
		User:     data.User,
		Project:  data.Project,
		Task:     data.Task,
		Billable: data.Billable,
		Start:    StringToData(strings.Join([]string{data.StartDate, data.StartTime}, " ")),
		End:      StringToData(strings.Join([]string{data.EndDate, data.EndTime}, " ")),
		Duration: StringToDuration(data.Duration),
	}
}
