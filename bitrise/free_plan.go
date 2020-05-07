package bitrise

const freePlanDefaultConcurrentBuildCount = 1
const freePlanDefaultMaximumBuildDurationInMinutes = 45
const freePlanDefaultMaximumBuildsPerMonth = 200
const freePlanDefaultMaximumTeamMembers = 2

type FreePlan struct {}

func (fp FreePlan) GetConcurrentBuildCount() int {
	return freePlanDefaultConcurrentBuildCount
}

func (fp FreePlan) GetMaximumBuildDurationInMinutes() int {
	return freePlanDefaultMaximumBuildDurationInMinutes
}

func (fp FreePlan) GetMaximumBuildsPerMonth() int {
	return freePlanDefaultMaximumBuildsPerMonth
}

func (fp FreePlan) GetMaximumTeamMembers() int {
	return freePlanDefaultMaximumTeamMembers
}
