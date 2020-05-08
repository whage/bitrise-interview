package bitrise

type PublicApp struct {
	App
	usesPublicAppLimits bool
}

func NewPublicApp(owner User) App {
	return App{
		owner: owner,
		usesPublicAppLimits: true,
		IPlan: NewPublicAppPlan(),
	}
}

func (a *App) OptOutFromPublicAppLimits() {
	a.usesPublicAppLimits = false
}

func (a *App) SetConcurrentBuildCount(n int) (bool, error) {
	plan := a.IPlan.(PublicAppPlan)
	ok, err := plan.SetConcurrentBuildCount(n)
	a.IPlan = plan
	return ok, err
}

func (a *App) SetMaximumBuildDurationInMinutes(n int) (bool, error) {
	plan := a.IPlan.(PublicAppPlan)
	ok, err := plan.SetMaximumBuildDurationInMinutes(n)
	a.IPlan = plan
	return ok, err
}
