package bitrise

const organizationPlanDefaultConcurrentBuildCount = 4
const organizationPlanDefaultMaximumBuildDurationInMinutes = 90

type OrganizationPlan struct {}

func (op OrganizationPlan) GetConcurrentBuildCount() int {
	return organizationPlanDefaultConcurrentBuildCount
}

func (op OrganizationPlan) GetMaximumBuildDurationInMinutes() int {
	return organizationPlanDefaultMaximumBuildDurationInMinutes
}

func (op OrganizationPlan) GetMaximumBuildsPerMonth() (bool, int) {
	return false, 0
}

func (op OrganizationPlan) GetMaximumTeamMembers() (bool, int) {
	return false, 0
}
