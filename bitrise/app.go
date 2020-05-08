package bitrise

type App struct {
	IPlan
	owner User
	isPublic bool // TODO: do we need this?
	usesPublicAppLimits bool
}

func (a App) GetPlan() IPlan {
	return a.IPlan
}

func NewApp(owner User, isPublic bool) App {
	app := App{
		owner: owner,
		isPublic: isPublic,
		usesPublicAppLimits: true,
	}
	if isPublic {
		app.IPlan = NewPublicAppPlan()
	}
	return app
}

// TODO: what if this app is not public?
func (a *App) OptOutFromPublicAppLimits() {
	a.usesPublicAppLimits = false
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
