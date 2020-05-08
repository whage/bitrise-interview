package bitrise

type PrivateApp struct {
	owner User
}

func NewPrivateApp(owner User) PrivateApp {
	return PrivateApp{
		owner: owner,
	}
}

func (a PrivateApp) GetOwner() User {
	return a.owner
}

func (a PrivateApp) GetConcurrentBuildCount() int {
	return a.owner.plan.GetConcurrentBuildCount()
}

func (a PrivateApp) GetMaximumBuildDurationInMinutes() int {
	return a.owner.plan.GetMaximumBuildDurationInMinutes()
}

func (a PrivateApp) GetMaximumBuildsPerMonth() (bool, int) {
	return a.owner.plan.GetMaximumBuildsPerMonth()
}

func (a PrivateApp) GetMaximumTeamMembers() (bool, int) {
	return a.owner.plan.GetMaximumTeamMembers()
}
