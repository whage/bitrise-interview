package main

import (
	"fmt"
	"github.com/whage/bitrise-interview/bitrise"
)

func main() {
	// 1. createa a user account on one of the plans
	user := bitrise.NewUser("dude@bitrise.io", bitrise.DeveloperPlan{})
	fmt.Println(user)

	// 2. create a public or private app
	publicApp := bitrise.NewApp(user, true)
	fmt.Println("new public app", publicApp)

	privateApp := bitrise.NewApp(user, false)
	fmt.Println("new private app", privateApp)

	// 3. set up custom limits for a public app
	plan, ok := publicApp.IPlan.(bitrise.PublicAppPlan) // TODO: this seems bad
	if ok {
		plan.SetConcurrentBuildCount(3)
		plan.SetMaximumBuildDurationInMinutes(60)
	}
	fmt.Println("plan after setting", plan)

	// 4. opt out from default public app limits
	// Note: "default" is incorrect in the spec I think
	publicApp.OptOutFromPublicAppLimits()
	fmt.Println("plan after opting out", plan)
	fmt.Println("publicApp.IPlan after opting out", publicApp.IPlan)

	// 5. get the limits of an app
	fmt.Println(publicApp.GetConcurrentBuildCount())
	fmt.Println(publicApp.GetMaximumBuildDurationInMinutes())
	fmt.Println(publicApp.GetMaximumBuildsPerMonth())
	fmt.Println(publicApp.GetMaximumTeamMembers())
}
