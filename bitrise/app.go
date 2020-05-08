package bitrise

import (
	"fmt"
)

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

// TODO: what if this app is not public?
func (a *App) OptOutFromPublicAppLimits() {
	a.usesPublicAppLimits = false
}

// TODO: what if this app is not public?
func (a *App) SetConcurrentBuildCount(n int) (bool, error) {
	if a.isPublic {
		plan := a.IPlan.(PublicAppPlan)
		ok, err := plan.SetConcurrentBuildCount(n)
		a.IPlan = plan
		return ok, err
	}
	return false, fmt.Errorf("Not a public app")
}

// TODO: what if this app is not public?
func (a *App) SetMaximumBuildDurationInMinutes(n int) (bool, error) {
	if a.isPublic {
		plan := a.IPlan.(PublicAppPlan)
		ok, err := plan.SetMaximumBuildDurationInMinutes(n)
		a.IPlan = plan
		return ok, err
	}
	return false, fmt.Errorf("Not a public app")
}
