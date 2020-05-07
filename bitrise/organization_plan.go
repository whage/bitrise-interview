package bitrise

type OrganizationPlan struct {}

func (op OrganizationPlan) GetConcurrentBuildCount() int {
	return 4
}

func (op OrganizationPlan) GetMaximumBuildDurationInMinutes() int {
	return 90
}

func (op OrganizationPlan) GetMaximumBuildsPerMonth() (bool, int) {
	return false, 0
}

func (op OrganizationPlan) GetMaximumTeamMembers() (bool, int) {
	return false, 0
}
