package main

type User struct {
	Id   string
	Name string
}

var users []*User

func GetUsers() []*User {
	return users
}

func SelectUser(id string) *User {
	for _, each := range users {
		if each.Id == id {
			return each
		}
	}
	return nil
}

func init() {
	users = append(users, &User{Id: "a01", Name: "Ben"})
	users = append(users, &User{Id: "a02", Name: "John"})
	users = append(users, &User{Id: "a03", Name: "Malik"})
}
