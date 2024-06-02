// database/user.go
package database

import (
	"database/sql"
	"fmt"

	"github.com/xiexiaohuozi/chatbot-system/model" // Replace with your actual module path
)

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, user *model.User) error {
	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last inserted ID: %w", err)
	}

	user.ID = int(lastInsertID)
	return nil
}

// GetUserByUsername fetches a user by their username
func GetUserByUsername(db *sql.DB, username string) (*model.User, error) {
	user := &model.User{}
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return user, nil
}