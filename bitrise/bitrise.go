package bitrise

type App struct {
	isPublic bool
	plan IPlan
}

func (a App) GetPlan() IPlan {
	return a.plan
}

type User struct {
	email string
	apps []Planner
}

type Planner interface {
	GetPlan() IPlan
}

type IPlan interface {
	GetConcurrentBuildCount() int
	GetMaximumBuildDurationInMinutes() int
	GetMaximumBuildsPerMonth() (bool, int)
	GetMaximumTeamMembers() (bool, int)
}