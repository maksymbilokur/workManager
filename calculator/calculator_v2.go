package calculator

import "github.com/workmanager/entity"

func UpdateMetricV2(data entity.InputData, project map[string]entity.ProjectMetric, teamMember map[string]entity.TeamMemberMetric) {
	// ProjectMetric
	if _, ok := project[data.Project]; ok {
		countOfUser := 0
		teamMembers := project[data.Project].TeamMembers
		for _, member := range project[data.Project].TeamMembers {
			if member == data.User {
				countOfUser++
			}
		}
		if countOfUser == 0 {
			teamMembers = append(teamMembers, data.User)
		}
		project[data.Project] = entity.ProjectMetric{
			WorkingHours: ProjectMetric[data.Project].WorkingHours + entity.StringToDuration(data.Duration),
			TeamMembers:  teamMembers,
		}
	} else {
		project[data.Project] = entity.ProjectMetric{
			WorkingHours: entity.StringToDuration(data.Duration),
			TeamMembers:  append([]string{}, data.User),
		}
	}

	//Member metric
	floatDuration := entity.StringToDuration(data.Duration)
	if _, ok := teamMember[data.User]; ok {
		teamMember[data.User] = entity.TeamMemberMetric{
			Salary:            teamMember[data.User].Salary + calculateOneSalary(floatDuration, data.Project),
			TotalWorkingHours: teamMember[data.User].TotalWorkingHours + floatDuration,
		}
	} else {
		teamMember[data.User] = entity.TeamMemberMetric{
			Salary:            calculateOneSalary(floatDuration, data.Project),
			TotalWorkingHours: floatDuration,
		}
	}

}
