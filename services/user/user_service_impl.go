package user

import (
	"bookmysalon/models"
	"bookmysalon/pkg/jwt"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct{}

func (us *UserServiceImpl) RegisterUser(db *sql.DB, u *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), BcryptCostFactor)
	if err != nil {
		return errors.New(HashingErrorMessage)
	}

	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;"
	if err := db.QueryRow(query, u.Username, hashedPassword).Scan(&u.ID); err != nil {
		return errors.New("failed to register user")
	}
	return nil
}

func (us *UserServiceImpl) LoginUser(db *sql.DB, u *models.User) (string, error) {
	var dbUser models.User
	query := "SELECT id, password FROM users WHERE username=$1;"
	if err := db.QueryRow(query, u.Username).Scan(&dbUser.ID, &dbUser.Password); err != nil {
		return "", errors.New(UserNotFoundMessage)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password)); err != nil {
		return "", errors.New(InvalidPasswordMessage)
	}

	token, err := jwt.GenerateToken(u.Username)
	if err != nil {
		return "", errors.New(TokenErrorMessage)
	}
	return token, nil
}

func (us *UserServiceImpl) FetchUserProfile(db *sql.DB, username string) (*models.User, error) {
	var u models.User
	query := "SELECT id, username FROM users WHERE username=$1;"
	if err := db.QueryRow(query, username).Scan(&u.ID, &u.Username); err != nil {
		return nil, errors.New(UserNotFoundMessage)
	}
	u.Password = "" // Ensure password is not exposed
	return &u, nil
}

func (us *UserServiceImpl) UpdateUserProfile(db *sql.DB, u *models.User) error {
	query := "UPDATE users SET email=$1, profile_image=$2 WHERE username=$3;"
	_, err := db.Exec(query, u.Email, u.ProfileImage, u.Username)
	return err
}

type PasswordChangeRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (us *UserServiceImpl) ChangeUserPassword(db *sql.DB, username, oldPassword, newPassword string) error {
	var currentHashedPassword string
	query := "SELECT password FROM users WHERE username=$1;"
	if err := db.QueryRow(query, username).Scan(&currentHashedPassword); err != nil {
		return errors.New(UserNotFoundMessage)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(currentHashedPassword), []byte(oldPassword)); err != nil {
		return errors.New(InvalidPasswordMessage)
	}

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), BcryptCostFactor)
	if err != nil {
		return errors.New(HashingErrorMessage)
	}

	updateQuery := "UPDATE users SET password=$1 WHERE username=$2;"
	_, err = db.Exec(updateQuery, newHashedPassword, username)
	return err
}

func (us *UserServiceImpl) DeleteUserAccount(db *sql.DB, username string) error {
	query := "DELETE FROM users WHERE username=$1;"
	_, err := db.Exec(query, username)
	return err
}
