package shared

type User struct {
	Id       string
	Email    string
	Username string
	Password string
}

func (u *User) Validate() bool {
	return u.Email != "" && u.Username != "" && u.Password != ""
}
