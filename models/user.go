package models

import (
	"errors"

	"github.com/yoshihyoda/golang_api/db"
	"github.com/yoshihyoda/golang_api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashString(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(retrievedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
