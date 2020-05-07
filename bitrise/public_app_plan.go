package bitrise

const publicAppPlanDefaultConcurrentBuildCount = 2
const publicAppPlanDefaultMaximumBuildDurationInMinutes = 45

// TODO: nothing protects these fields from being changed directly
type PublicAppPlan struct {
	concurrentBuildCount int
	maximumBuildDurationInMinutes int
}

func (pap PublicAppPlan) GetConcurrentBuildCount() int {
	if pap.concurrentBuildCount != 0 {
		return pap.concurrentBuildCount
	}
	return publicAppPlanDefaultConcurrentBuildCount
}

func (pap PublicAppPlan) GetMaximumBuildDurationInMinutes() int {
	if pap.maximumBuildDurationInMinutes != 0 {
		return pap.maximumBuildDurationInMinutes
	}
	return publicAppPlanDefaultMaximumBuildDurationInMinutes
}

func (pap PublicAppPlan) GetMaximumBuildsPerMonth() (bool, int) {
	return false, 0
}

func (pap PublicAppPlan) GetMaximumTeamMembers() (bool, int) {
	return false, 0
}

func (pap *PublicAppPlan) SetConcurrentBuildCount(n int) {
	if n > publicAppPlanDefaultConcurrentBuildCount {
		pap.concurrentBuildCount = n
	}
}

func (pap *PublicAppPlan) SetMaximumBuildDurationInMinutes(n int) {
	if n > publicAppPlanDefaultMaximumBuildDurationInMinutes {
		pap.maximumBuildDurationInMinutes = n
	}
}
