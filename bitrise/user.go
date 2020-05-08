package bitrise

type User struct {
	email string
	plan IPlan
	apps []IApp
}

func NewUser(email string, plan IPlan) User {
	return User{
		email: email,
		plan: plan,
	}
}

func (u *User) AddApp(app IApp) {
	u.apps = append(u.apps, app)
}
