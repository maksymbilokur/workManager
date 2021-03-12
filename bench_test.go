package main

import (
	"testing"

	"github.com/workmanager/calculator"
	"github.com/workmanager/entity"

	"github.com/workmanager/parser"
)

func BenchmarkRearerWithmetric(b *testing.B) {

	for n := 0; n < b.N; n++ {
		calculator.ProjectMetric = make(map[string]entity.ProjectMetric)
		calculator.MemberMetric = make(map[string]entity.TeamMemberMetric)
		parser.ReaderWithoutMetric("sanitaze.csv", 14)
	}
}

func BenchmarkReader(b *testing.B) {

	for n := 0; n < b.N; n++ {
		var activities []entity.ActivityData
		inputActivities := parser.Reader("sanitaze.csv", 14)
		for _, d := range inputActivities {
			activities = append(activities, entity.InputToActivity(d))
		}
		calculator.ProjectMetric = make(map[string]entity.ProjectMetric)
		calculator.MemberMetric = make(map[string]entity.TeamMemberMetric)
		calculator.WorkingHoursAndUsersForProjects(activities)
		calculator.SalaryHoursForAllMembers(activities)
	}
}
