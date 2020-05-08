package bitrise

type App struct {
	IPlan
	owner User
	usesPublicAppLimits bool
}

func (a App) GetPlan() IPlan {
	return a.IPlan
}

func NewApp(owner User) App {
	return App{
		owner: owner,
	}
}

func (a App) GetConcurrentBuildCount() int {
	if !a.usesPublicAppLimits {
		a.owner.plan.GetConcurrentBuildCount()
	}
	return a.IPlan.GetConcurrentBuildCount()
}

func (a App) GetMaximumBuildDurationInMinutes() int {
	if !a.usesPublicAppLimits {
		a.owner.plan.GetMaximumBuildDurationInMinutes()
	}
	return a.IPlan.GetMaximumBuildDurationInMinutes()
}

func (a App) GetMaximumBuildsPerMonth() (bool, int) {
	if !a.usesPublicAppLimits {
		a.owner.plan.GetMaximumBuildsPerMonth()
	}
	return a.IPlan.GetMaximumBuildsPerMonth()
}

func (a App) GetMaximumTeamMembers() (bool, int) {
	if !a.usesPublicAppLimits {
		a.owner.plan.GetMaximumTeamMembers()
	}
	return a.IPlan.GetMaximumTeamMembers()
}
