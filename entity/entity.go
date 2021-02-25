package entity

import (
	"strconv"
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

//timeparce
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
	//TODO time.ParseDuration
	sepd := strings.Split(d, ":")
	hour, _ := strconv.Atoi(sepd[0])
	min, _ := strconv.Atoi(sepd[1])
	sec, _ := strconv.Atoi(sepd[2])

	return float32(hour) + float32(min)/60 + float32(sec)/3600
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
