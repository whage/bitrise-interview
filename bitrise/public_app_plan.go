package bitrise

import (
	"fmt"
)

const publicAppPlanDefaultConcurrentBuildCount = 2
const publicAppPlanDefaultMaximumBuildDurationInMinutes = 45

// TODO: nothing protects these fields from being changed directly
type PublicAppPlan struct {
	concurrentBuildCount int
	maximumBuildDurationInMinutes int
}

func NewPublicAppPlan() PublicAppPlan {
	return PublicAppPlan{
		concurrentBuildCount: publicAppPlanDefaultConcurrentBuildCount,
		maximumBuildDurationInMinutes: publicAppPlanDefaultMaximumBuildDurationInMinutes,
	}
}

func (pap PublicAppPlan) GetConcurrentBuildCount() int {
	return pap.concurrentBuildCount
}

func (pap PublicAppPlan) GetMaximumBuildDurationInMinutes() int {
	return pap.maximumBuildDurationInMinutes
}

func (pap PublicAppPlan) GetMaximumBuildsPerMonth() (bool, int) {
	return false, 0
}

func (pap PublicAppPlan) GetMaximumTeamMembers() (bool, int) {
	return false, 0
}

func (pap *PublicAppPlan) SetConcurrentBuildCount(n int) (bool, error) {
	if n >= publicAppPlanDefaultConcurrentBuildCount {
		pap.concurrentBuildCount = n
		return true, nil
	}
	return false, fmt.Errorf("New value (%d) is less than plan default (%d)", n, publicAppPlanDefaultConcurrentBuildCount)
}

func (pap *PublicAppPlan) SetMaximumBuildDurationInMinutes(n int) (bool, error) {
	if n >= publicAppPlanDefaultMaximumBuildDurationInMinutes {
		pap.maximumBuildDurationInMinutes = n
		return true, nil
	}
	return false, fmt.Errorf("New value (%d) is less than plan default (%d)", n, publicAppPlanDefaultMaximumBuildDurationInMinutes)
}
