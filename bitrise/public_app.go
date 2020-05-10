package bitrise

type PublicApp struct {
	PrivateApp
	PublicAppPlan
	usesPublicAppLimits bool
}

func NewPublicApp(base PrivateApp, plan PublicAppPlan) PublicApp {
	return PublicApp{
		PrivateApp: base,
		PublicAppPlan: plan,
		usesPublicAppLimits: true,
	}
}

func (pa *PublicApp) OptOutFromPublicAppLimits() {
	pa.usesPublicAppLimits = false
}

func (pa PublicApp) GetConcurrentBuildCount() int {
	if !pa.usesPublicAppLimits {
		return pa.owner.plan.GetConcurrentBuildCount()
	}
	return pa.PublicAppPlan.GetConcurrentBuildCount()
}

func (pa PublicApp) GetMaximumBuildDurationInMinutes() int {
	if !pa.usesPublicAppLimits {
		return pa.owner.plan.GetMaximumBuildDurationInMinutes()
	}
	return pa.PublicAppPlan.GetMaximumBuildDurationInMinutes()
}

func (pa PublicApp) GetMaximumBuildsPerMonth() (bool, int) {
	if !pa.usesPublicAppLimits {
		return pa.owner.plan.GetMaximumBuildsPerMonth()
	}
	return pa.PublicAppPlan.GetMaximumBuildsPerMonth()
}

func (pa PublicApp) GetMaximumTeamMembers() (bool, int) {
	if !pa.usesPublicAppLimits {
		return pa.owner.plan.GetMaximumTeamMembers()
	}
	return pa.PublicAppPlan.GetMaximumTeamMembers()
}
