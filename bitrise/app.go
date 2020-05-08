package bitrise

type App struct {
	owner User
}

func NewApp(owner User) App {
	return App{
		owner: owner,
	}
}

func (a App) GetConcurrentBuildCount() int {
	return a.owner.plan.GetConcurrentBuildCount()
}

func (a App) GetMaximumBuildDurationInMinutes() int {
	return a.owner.plan.GetMaximumBuildDurationInMinutes()
}

func (a App) GetMaximumBuildsPerMonth() (bool, int) {
	return a.owner.plan.GetMaximumBuildsPerMonth()
}

func (a App) GetMaximumTeamMembers() (bool, int) {
	return a.owner.plan.GetMaximumTeamMembers()
}
