package bitrise

const developerPlanDefaultConcurrentBuildCount = 2
const developerPlanDefaultMaximumBuildDurationInMinutes = 45

type DeveloperPlan struct {}

func (dp DeveloperPlan) GetConcurrentBuildCount() int {
	return developerPlanDefaultConcurrentBuildCount
}

func (dp DeveloperPlan) GetMaximumBuildDurationInMinutes() int {
	return developerPlanDefaultMaximumBuildDurationInMinutes
}

func (dp DeveloperPlan) GetMaximumBuildsPerMonth() (bool, int) {
	return false, 0
}

func (dp DeveloperPlan) GetMaximumTeamMembers() (bool, int) {
	return false, 0
}
