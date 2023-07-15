package user

import "fmt"

type User struct {
	Username string
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("username: %s, password: %s\n", u.Username, u.Password)
}
