package models

import (
	"database/sql"
	"event_booking/utils"
	"fmt"
)

type User struct {
	Id       int
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (user *User) SignUp(DB *sql.DB) error {
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id;
	`
	stmt, err := DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (user *User) ValidateCredential(DB *sql.DB) (bool, error) {
	query := `
		select id, password
		from users
		where email = $1
	`
	var retrivedPassword string
	err := DB.QueryRow(query, user.Email).Scan(&user.Id, &retrivedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	isPasswordValid := utils.CheckPassword(user.Password, retrivedPassword)
	return isPasswordValid, nil

}
