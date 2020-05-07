package bitrise

type User struct {
	email string
	plan IPlan
	apps []IPlanOwner
}

func (u User) GetPlan() IPlan {
	return u.plan
}

func NewUser(email string, plan IPlan) User {
	return User{
		email: email,
		plan: plan,
	}
}
