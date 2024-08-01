package models

import (
	"errors"

	"osho.com/db"
	"osho.com/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
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

func (u *User) Authenticate() error{
     query := "Select id,password from users where email = ?"
	 row := db.DB.QueryRow(query, u.Email)
	 var retrievedPassword string
	 err :=row.Scan(&u.ID,&retrievedPassword)
	 if err != nil {
		 return errors.New("invalid Credentials")
	 }
	 passwordIsValid := utils.ComparePassword(retrievedPassword, u.Password)
	 if !passwordIsValid {
		 return errors.New("invalid Credentials")
	 }
	 return nil
}
