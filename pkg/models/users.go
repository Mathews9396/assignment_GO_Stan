package models

import (
	"fmt"
)

type User struct {
    ID        string `json:"id"`
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
    Email     string `json:"email"`
}
var users  = []User {
	{ID: "1", Firstname: "John", Lastname: "Doe", Email: "john@example.com"},
    {ID: "2", Firstname: "Jane", Lastname: "Smith", Email: "jane@example.com"},
}

func GetAllUsers() []User{
	return users
}

func CreateUser(user *User) User{
	users  = append(users, *user)
	return *user
}

func GetUser(Id string) *User{
	for _, user := range users{
		if user.ID == Id {
			return &user
		}
	}
	return nil
}

func UpdateUser(user *User) *User{
	for index, item := range users {
		if item.ID == user.ID{
			fmt.Printf(`The user has been found and now deleting and adding new data`)

			users = append(users[:index], users[index+1:]...)
			users = append(users, *user)
			return user
		}
	}
	return nil
}

func DeleteUser(ID string) bool {
	for index, item := range users {
		if item.ID == ID{
			fmt.Printf(`The user has been found and now deleting user`)
			users = append(users[:index], users[index+1:]...)
			return true
		}
	}
	return false
}