package main

import (
	"fmt"
	"github.com/whage/bitrise-interview/bitrise"
)

func main() {
	// 1. createa a user account on one of the plans
	user := bitrise.NewUser("dude@bitrise.io", bitrise.OrganizationPlan{})
	fmt.Println("New user:", user)

	// 2. create a public or private app
	publicApp := bitrise.NewPublicApp(user)
	user.AddApp(publicApp)
	fmt.Println("New public app:", publicApp)

	privateApp := bitrise.NewPrivateApp(user)
	user.AddApp(publicApp)
	fmt.Println("New private app:", privateApp)

	// 3. set up custom limits for a public app
	publicApp.SetConcurrentBuildCount(3)
	publicApp.SetMaximumBuildDurationInMinutes(60)
	fmt.Println("publicApp after setting custom limits:", publicApp)

	// Note: setting custom limits on a private app is a compile time error:
	//privateApp.SetConcurrentBuildCount(3)

	// 4. opt out from default public app limits
	// Note: "default" is incorrect in the spec I think
	publicApp.OptOutFromPublicAppLimits()
	fmt.Println("publicApp after opting out:", publicApp)

	// 5. get the limits of opted-out public app
	fmt.Println(publicApp.GetConcurrentBuildCount())
	fmt.Println(publicApp.GetMaximumBuildDurationInMinutes())
	fmt.Println(publicApp.GetMaximumBuildsPerMonth())
	fmt.Println(publicApp.GetMaximumTeamMembers())
}
