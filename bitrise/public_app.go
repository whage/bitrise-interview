package bitrise

type PublicApp struct {
	App
	PublicAppPlan
	usesPublicAppLimits bool
}

func NewPublicApp(owner User) PublicApp {
	return PublicApp{
		App: App{
			owner: owner,
		},
		PublicAppPlan: NewPublicAppPlan(),
		usesPublicAppLimits: true,
	}
}

func (a *PublicApp) OptOutFromPublicAppLimits() {
	a.usesPublicAppLimits = false
}

func (a *PublicApp) SetConcurrentBuildCount(n int) (bool, error) {
	return a.PublicAppPlan.SetConcurrentBuildCount(n)
}

func (a *PublicApp) SetMaximumBuildDurationInMinutes(n int) (bool, error) {
	return a.PublicAppPlan.SetMaximumBuildDurationInMinutes(n)
}

func (a PublicApp) GetConcurrentBuildCount() int {
	if !a.usesPublicAppLimits {
		a.owner.plan.GetConcurrentBuildCount()
	}
	return a.PublicAppPlan.GetConcurrentBuildCount()
}

func (a PublicApp) GetMaximumBuildDurationInMinutes() int {
	if !a.usesPublicAppLimits {
		a.owner.plan.GetMaximumBuildDurationInMinutes()
	}
	return a.PublicAppPlan.GetMaximumBuildDurationInMinutes()
}

func (a PublicApp) GetMaximumBuildsPerMonth() (bool, int) {
	if !a.usesPublicAppLimits {
		a.owner.plan.GetMaximumBuildsPerMonth()
	}
	return a.PublicAppPlan.GetMaximumBuildsPerMonth()
}

func (a PublicApp) GetMaximumTeamMembers() (bool, int) {
	if !a.usesPublicAppLimits {
		a.owner.plan.GetMaximumTeamMembers()
	}
	return a.PublicAppPlan.GetMaximumTeamMembers()
}
