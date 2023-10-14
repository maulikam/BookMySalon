package user

import (
	"bookmysalon/models"
	"database/sql"
)

type UserService interface {
	RegisterUser(db *sql.DB, u *models.User) error
	LoginUser(db *sql.DB, u *models.User) (string, error)
	FetchUserProfile(db *sql.DB, username string) (*models.User, error)
	UpdateUserProfile(db *sql.DB, u *models.User) error
	ChangeUserPassword(db *sql.DB, username, oldPassword, newPassword string) error
	DeleteUserAccount(db *sql.DB, username string) error
}
