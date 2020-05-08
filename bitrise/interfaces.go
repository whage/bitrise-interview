package bitrise

type IApp interface {
	GetOwner() User
}

type IPlan interface {
	GetConcurrentBuildCount() int
	GetMaximumBuildDurationInMinutes() int
	GetMaximumBuildsPerMonth() (bool, int)
	GetMaximumTeamMembers() (bool, int)
}
