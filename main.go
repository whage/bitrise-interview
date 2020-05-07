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
	publicApp := bitrise.NewApp(true)
	fmt.Println(publicApp)

	// 3. set up custom limits for a public app
	plan, ok := publicApp.IPlan.(bitrise.PublicAppPlan) // TODO: this seems bad
	if ok {
		plan.SetConcurrentBuildCount(5)
	}
}
