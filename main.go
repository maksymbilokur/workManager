package main

import (
	"fmt"
	"taskmanager/entity"
	"taskmanager/parser"
	"time"
)

var salaryList map[string]entity.ProjectSalary
var project map[string]entity.ProjectMetric

const RegularHours = 8.0

func main() {
	list := parser.Reader("sanitaze.csv", 14)
	var activity []entity.ActivityData
	for _, l := range list {
		activity = append(activity, entity.InputToActivity(l))
	}
	initBaseMap()

	/*
		cSalary := calculateSalaryForUser("Princess Jailyn Weissnat", activity, time.Time{}, time.Time{})
		fmt.Println("in all project salary:",cSalary)
		oneProjectSalary := calculateSalaryForUserInProject("Princess Jailyn Weissnat","hic culpa", activity, time.Time{}, time.Time{})
		fmt.Println("in one project salary:",oneProjectSalary)
		workingHoursAndUsersForProjects(activity)
		for k, v := range project {
			fmt.Println(k, v)
		}
	*/

	memberStats := SalaryHoursForMember("Princess Jailyn Weissnat", activity, time.Time{}, time.Time{})
	fmt.Println("stats for member:", memberStats)
}

//prepare for input date range
func calculateSalaryForUser(user string, activities []entity.ActivityData, _ time.Time, _ time.Time) float32 {
	var salary float32
	for _, a := range activities {
		if a.User == user /*&& a.Billable != "No"*/ {
			if a.Duration <= RegularHours {
				salary += salaryList[a.Project].Salary * a.Duration
			} else {
				salary += salaryList[a.Project].Salary*RegularHours + salaryList[a.Project].OvertimeSalary*(a.Duration-RegularHours)
			}
		}
	}
	return salary
}

func calculateSalaryForUserInProject(user, project string, activities []entity.ActivityData, _ time.Time, _ time.Time) float32 {
	var salary float32
	for _, a := range activities {
		if a.User == user && a.Project == project /*&& a.Billable != "No"*/ {
			if a.Duration <= RegularHours {
				salary += salaryList[a.Project].Salary * a.Duration
			} else {
				salary += salaryList[a.Project].Salary*RegularHours + salaryList[a.Project].OvertimeSalary*(a.Duration-RegularHours)
			}
		}
	}
	return salary
}

func workingHoursAndUsersForProjects(activities []entity.ActivityData) {

	project = make(map[string]entity.ProjectMetric)
	for _, a := range activities {
		if _, ok := project[a.Project]; ok {
			countOfUser := 0
			teamMembers := project[a.Project].TeamMembers
			for _, member := range project[a.Project].TeamMembers {
				if member == a.User {
					countOfUser++
				}
			}
			if countOfUser == 0 {
				teamMembers = append(teamMembers, a.User)
			}
			project[a.Project] = entity.ProjectMetric{
				WorkingHours: project[a.Project].WorkingHours + a.Duration,
				TeamMembers:  teamMembers,
			}
		} else {
			project[a.Project] = entity.ProjectMetric{
				WorkingHours: a.Duration,
				TeamMembers:  append([]string{}, a.User),
			}
		}
	}
}

func SalaryHoursForMember(user string, activities []entity.ActivityData, _ time.Time, _ time.Time) entity.TeamMemberMetric {
	//salary, total working hours
	total := entity.TeamMemberMetric{}

	for _, a := range activities {
		if a.User == user /*&& a.Billable != "No"*/ {
			if a.Duration <= RegularHours {
				total.Salary += salaryList[a.Project].Salary * a.Duration
			} else {
				total.Salary += salaryList[a.Project].Salary*RegularHours + salaryList[a.Project].OvertimeSalary*(a.Duration-RegularHours)
			}
			fmt.Println(total.TotalWorkingHours, a.Duration)
			total.TotalWorkingHours += a.Duration
		}
	}
	return total
}

//temporary method
func initBaseMap() {
	salaryList = make(map[string]entity.ProjectSalary)
	type args struct {
		name   string
		salary entity.ProjectSalary
	}
	fields := []args{
		{
			name:   "quia autem",
			salary: entity.ProjectSalary{Salary: 1.0, OvertimeSalary: 2.0},
		},
		{
			name:   "omnis doloribus",
			salary: entity.ProjectSalary{Salary: 1.0, OvertimeSalary: 2.0},
		},
		{
			name:   "id ex",
			salary: entity.ProjectSalary{Salary: 1.0, OvertimeSalary: 2.0},
		},
		{
			name:   "mollitia minus",
			salary: entity.ProjectSalary{Salary: 1.0, OvertimeSalary: 2.0},
		},
		{
			name:   "hic culpa",
			salary: entity.ProjectSalary{Salary: 1.0, OvertimeSalary: 2.0},
		},
	}
	for _, f := range fields {
		salaryList[f.name] = entity.ProjectSalary{
			Salary:         f.salary.Salary,
			OvertimeSalary: f.salary.OvertimeSalary,
		}
	}

}
