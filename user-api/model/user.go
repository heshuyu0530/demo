package model

type User struct {
    UserName string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]string)

func Register(user User) bool{
	if _,exists := users[user.UserName]; exists {
		return false
	}
	users[user.UserName] = user.Password
	return true
}

func Login(user User) bool{
	if password,exists := users[user.UserName]; exists && password == user.Password {
		return true
	}
	return false
}
