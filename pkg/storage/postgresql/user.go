package postgresql

import (
	"database/sql"

	"github.com/tbaud0n/sample-api-architecture"
	"github.com/tbaud0n/sample-api-architecture/pkg/logger"
)

// UserService is a mock of the api.UserService
type UserService struct {
	DB *sql.DB
}

// Get returns the user from its id
func (us *UserService) Get(id int64) (*api.User, error) {
	logger.LogDebug("NOT IMPLEMENTED")
	return nil, nil
}

// Query returns the users matching the queryFilter given as argument
func (us *UserService) Query(f interface{}) ([]*api.User, error) {
	qf := f.(*QueryFilter)

	s := SQLQuery(`SELECT id, lastname, firstname FROM "user"`, qf)

	rows, err := us.DB.Query(s, qf.Args...)
	if err != nil {
		return nil, logger.LogError(err)
	}

	defer func() {
		logger.LogError(rows.Close())
	}()

	users := []*api.User{}

	for rows.Next() {
		user := &api.User{}
		err := rows.Scan(&user.ID, &user.Lastname, &user.Firstname)
		if err != nil {
			return nil, logger.LogError(err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, logger.LogError(err)
	}

	return users, nil
}

// Create inserts the new user to the DB
func (us *UserService) Create(u *api.User) error {
	logger.LogDebug("NOT IMPLEMENTED")
	return nil
}

// Update the given user in the DB
func (us *UserService) Update(u *api.User) error {
	logger.LogDebug("NOT IMPLEMENTED")
	return nil
}

// Delete the user with the given id from the DB
func (us *UserService) Delete(id int64) error {
	logger.LogDebug("NOT IMPLEMENTED")
	return nil
}
