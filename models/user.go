package models

import (
	"errors"
	"example/rest_api/db"
	"example/rest_api/utils"
	"log"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO USERS (email,password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return nil
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	user.ID = userId

	return nil
}

func (user User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row :=
		db.DB.QueryRow(query, user.Email)
	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	log.Println("User Email: ", user.Email)
	log.Println("Retrieved Password (Before Scan): ", retrievedPassword)

	if err != nil {
		log.Println("ERRROR: ", err)
		return errors.New("invalid cradentials")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid cradentials")
	}

	log.Println("password is valid")

	return nil
}
