package bitrise

type App struct {
	isPublic bool // TODO: do we need this?
	IPlan
}

func (a App) GetPlan() IPlan {
	return a.IPlan
}

func NewApp(isPublic bool) App {
	app := App{
		isPublic: isPublic,
	}
	if isPublic {
		app.IPlan = PublicAppPlan{}
	}
	return app
}
