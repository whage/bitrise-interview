package bitrise

type DeveloperPlan struct {}

func (dp DeveloperPlan) GetConcurrentBuildCount() int {
	return 2
}

func (dp DeveloperPlan) GetMaximumBuildDurationInMinutes() int {
	return 45
}

func (dp DeveloperPlan) GetMaximumBuildsPerMonth() (bool, int) {
	return false, 0
}

func (dp DeveloperPlan) GetMaximumTeamMembers() (bool, int) {
	return false, 0
}
