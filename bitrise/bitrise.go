package bitrise

type IPlanOwner interface {
	GetPlan() IPlan
}

type IPlan interface {
	GetConcurrentBuildCount() int
	GetMaximumBuildDurationInMinutes() int
	GetMaximumBuildsPerMonth() (bool, int)
	GetMaximumTeamMembers() (bool, int)
}
