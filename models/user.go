package models

import (
	"errors"
	"fmt"
)

// User details
type User struct {
	ID    int
	FName string
	LName string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers Get all users
func GetUsers() []*User {
	return users
}

// AddUser Add a new user
func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("new users must not include id or it must be set to zero")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

// UpdateUser Update exisitng user
func UpdateUser(user User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == user.ID {
			users[i] = &user
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user with id '%v' not found", user.ID)
}

// RemoveUserByID Removes user based in ID
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with id '%v' not found", id)
}

// GetUserByID Get users by ID
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with id '%v' not found", id)
}
