package calculator

import (
	"github.com/workmanager/entity"
)

const RegularHours = 8.0

var (
	ProjectMetric map[string]entity.ProjectMetric
	MemberMetric  map[string]entity.TeamMemberMetric
	Projects      map[string]entity.ProjectSalary
)

func init() {
	ProjectMetric = make(map[string]entity.ProjectMetric)
	MemberMetric = make(map[string]entity.TeamMemberMetric)
	Projects = make(map[string]entity.ProjectSalary)
}

func calculateOneSalary(duration float32, projectName string) float32 {
	if duration <= RegularHours {
		return Projects[projectName].Salary * duration
	} else {
		return Projects[projectName].Salary*RegularHours + Projects[projectName].OvertimeSalary*(duration-RegularHours)
	}
}

//prepare for input date range
func CalculateSalaryForUser(user string, activities []entity.ActivityData) float32 {
	var salary float32
	for _, a := range activities {
		if a.User == user {
			salary += calculateOneSalary(a.Duration, a.Project)
		}
	}
	return salary
}

func CalculateSalaryForUserInProject(user, project string, activities []entity.ActivityData) float32 {
	var salary float32
	for _, a := range activities {
		if a.User == user && a.Project == project {
			salary += calculateOneSalary(a.Duration, a.Project)
		}
	}
	return salary
}

func WorkingHoursAndUsersForProjects(activities []entity.ActivityData) {

	ProjectMetric = make(map[string]entity.ProjectMetric)
	for _, a := range activities {
		if _, ok := ProjectMetric[a.Project]; ok {
			countOfUser := 0
			teamMembers := ProjectMetric[a.Project].TeamMembers
			for _, member := range ProjectMetric[a.Project].TeamMembers {
				if member == a.User {
					countOfUser++
				}
			}
			if countOfUser == 0 {
				teamMembers = append(teamMembers, a.User)
			}
			ProjectMetric[a.Project] = entity.ProjectMetric{
				WorkingHours: ProjectMetric[a.Project].WorkingHours + a.Duration,
				TeamMembers:  teamMembers,
			}
		} else {
			ProjectMetric[a.Project] = entity.ProjectMetric{
				WorkingHours: a.Duration,
				TeamMembers:  append([]string{}, a.User),
			}
		}
	}
}

func SalaryHoursForMember(user string, activities []entity.ActivityData) entity.TeamMemberMetric {
	//salary, total working hours
	total := entity.TeamMemberMetric{}

	for _, a := range activities {
		if a.User == user /*&& a.Billable != "No"*/ {
			total.Salary += calculateOneSalary(a.Duration, a.Project)
			total.TotalWorkingHours += a.Duration
		}
	}
	return total
}

func SalaryHoursForAllMembers(activities []entity.ActivityData) {
	MemberMetric = make(map[string]entity.TeamMemberMetric)
	for _, a := range activities {
		if _, ok := MemberMetric[a.User]; !ok {
			MemberMetric[a.User] = SalaryHoursForMember(a.User, activities)
		}
	}
}

func UpdateMetric(data entity.InputData) {
	// ProjectMetric
	if _, ok := ProjectMetric[data.Project]; ok {
		countOfUser := 0
		teamMembers := ProjectMetric[data.Project].TeamMembers
		for _, member := range ProjectMetric[data.Project].TeamMembers {
			if member == data.User {
				countOfUser++
			}
		}
		if countOfUser == 0 {
			teamMembers = append(teamMembers, data.User)
		}
		ProjectMetric[data.Project] = entity.ProjectMetric{
			WorkingHours: ProjectMetric[data.Project].WorkingHours + entity.StringToDuration(data.Duration),
			TeamMembers:  teamMembers,
		}
	} else {
		ProjectMetric[data.Project] = entity.ProjectMetric{
			WorkingHours: entity.StringToDuration(data.Duration),
			TeamMembers:  append([]string{}, data.User),
		}
	}

	//Member metric
	floatDuration := entity.StringToDuration(data.Duration)
	if _, ok := ProjectMetric[data.Project]; ok {
		MemberMetric[data.Project] = entity.TeamMemberMetric{
			Salary:            MemberMetric[data.Project].Salary + calculateOneSalary(floatDuration, data.Project),
			TotalWorkingHours: MemberMetric[data.Project].TotalWorkingHours + floatDuration,
		}
	} else {
		MemberMetric[data.Project] = entity.TeamMemberMetric{
			Salary:            calculateOneSalary(floatDuration, data.Project),
			TotalWorkingHours: floatDuration,
		}
	}

}
