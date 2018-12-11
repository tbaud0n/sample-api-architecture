package api

import "net/http"

// Here we define the structs and the

//User defines a user
type User struct {
	ID        int64
	Firstname string
	Lastname  string
}

// UserStorageService describes the methods required to handle the storage of the users
type UserStorageService interface {
	Get(id int64) (*User, error)
	Query(queryFilter interface{}) ([]*User, error)
	Create(u *User) error
	Update(u *User) error
	Delete(id int64) error
}

// QueryFilterFromHTTPRequest is the method called by the HTTP handler
// to generate the queryFilter passed to the query func of each storage service
type QueryFilterFromHTTPRequest func(r *http.Request) (interface{}, error)
