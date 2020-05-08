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

func (pa *PublicApp) OptOutFromPublicAppLimits() {
	pa.usesPublicAppLimits = false
}

func (pa PublicApp) GetConcurrentBuildCount() int {
	if !pa.usesPublicAppLimits {
		pa.owner.plan.GetConcurrentBuildCount()
	}
	return pa.PublicAppPlan.GetConcurrentBuildCount()
}

func (pa PublicApp) GetMaximumBuildDurationInMinutes() int {
	if !pa.usesPublicAppLimits {
		pa.owner.plan.GetMaximumBuildDurationInMinutes()
	}
	return pa.PublicAppPlan.GetMaximumBuildDurationInMinutes()
}

func (pa PublicApp) GetMaximumBuildsPerMonth() (bool, int) {
	if !pa.usesPublicAppLimits {
		pa.owner.plan.GetMaximumBuildsPerMonth()
	}
	return pa.PublicAppPlan.GetMaximumBuildsPerMonth()
}

func (pa PublicApp) GetMaximumTeamMembers() (bool, int) {
	if !pa.usesPublicAppLimits {
		pa.owner.plan.GetMaximumTeamMembers()
	}
	return pa.PublicAppPlan.GetMaximumTeamMembers()
}
