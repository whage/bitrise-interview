package bitrise

type PublicApp struct {
	App
	usesPublicAppLimits bool
}

func NewPublicApp(owner User) PublicApp {
	return PublicApp{
		App: App{
			owner: owner,
			IPlan: NewPublicAppPlan(),
		},
		usesPublicAppLimits: true,
	}
}

func (a *PublicApp) OptOutFromPublicAppLimits() {
	a.usesPublicAppLimits = false
}

func (a *PublicApp) SetConcurrentBuildCount(n int) (bool, error) {
	plan := a.IPlan.(PublicAppPlan)
	ok, err := plan.SetConcurrentBuildCount(n)
	a.IPlan = plan
	return ok, err
}

func (a *PublicApp) SetMaximumBuildDurationInMinutes(n int) (bool, error) {
	plan := a.IPlan.(PublicAppPlan)
	ok, err := plan.SetMaximumBuildDurationInMinutes(n)
	a.IPlan = plan
	return ok, err
}
