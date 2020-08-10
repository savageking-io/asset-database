package main

import (
	"net/http"
)

type User struct {
}

// Get returns information about the user
func (u *User) Get(w http.ResponseWriter, r *http.Request) {

}

// Login logs user in
func (u *User) Login(w http.ResponseWriter, r *http.Request) {

}

// Logout ends user session
func (u *User) Logout(w http.ResponseWriter, r *http.Request) {

}

// Create creates new user
func (u *User) Create(w http.ResponseWriter, r *http.Request) {

}

// Delete deletes a user
func (u *User) Delete(w http.ResponseWriter, r *http.Request) {

}
