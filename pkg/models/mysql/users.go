package mysql

import (
	"database/sql"

	"cubeguerrero.com/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

// Insert method to add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate method to verify whether a user exists with the
// provided email address and password. Returns the relevant
// user ID
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get method to fetch details for a specific user based
// on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
