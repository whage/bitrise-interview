package bitrise

type PublicAppPlan struct {
	concurrentBuildCount int
	maximumBuildDurationInMinutes int
}

func (pap PublicAppPlan) GetConcurrentBuildCount() int {
	if pap.concurrentBuildCount != 0 {
		return pap.concurrentBuildCount
	}
	return 2
}

func (pap PublicAppPlan) GetMaximumBuildDurationInMinutes() int {
	if pap.maximumBuildDurationInMinutes != 0 {
		return pap.maximumBuildDurationInMinutes
	}
	return 45
}

func (pap PublicAppPlan) GetMaximumBuildsPerMonth() (bool, int) {
	return false, 0
}

func (pap PublicAppPlan) GetMaximumTeamMembers() (bool, int) {
	return false, 0
}

func (pap *PublicAppPlan) SetConcurrentBuildCount(n int) {
	pap.concurrentBuildCount = n
}

func (pap *PublicAppPlan) SetMaximumBuildDurationInMinutes(n int) {
	pap.maximumBuildDurationInMinutes = n
}
