package bitrise

type FreePlan struct {}

func (fp FreePlan) GetConcurrentBuildCount() int {
	return 1
}

func (fp FreePlan) GetMaximumBuildDurationInMinutes() int {
	return 10
}

func (fp FreePlan) GetMaximumBuildsPerMonth() int {
	return 200
}

func (fp FreePlan) GetMaximumTeamMembers() int {
	return 2
}
