package users

import (
	"fmt"
	"github.com/StarCloud2/StarCloud2-api/datasources/mysql/users_db"
	"github.com/StarCloud2/StarCloud2-api/utils/date_utils"
	"github.com/StarCloud2/StarCloud2-api/utils/errors"
	"strings"
)

/*
	INFO:	Data Access Object
			Single point of database connection
*/

/*
	INFO: placeholders vor different drivers
MySQL               PostgreSQL            Oracle
=====               ==========            ======
WHERE col = ?       WHERE col = $1        WHERE col = :col
VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
*/

const (
	indexUniqueEmail = "email_UNIQUE"
	indexNoRows      = "no rows in result set"
	queryIndertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES( ?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser  = "DELETE FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInterlanServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), indexNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("user: %d not found", user.Id))
		}
		fmt.Println(err)
		return errors.NewInterlanServerError(
			fmt.Sprintf("error when trying to get user: %d", user.Id))
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryIndertUser)
	if err != nil {
		return errors.NewInterlanServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(
				fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInterlanServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInterlanServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInterlanServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInterlanServerError(err.Error())
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInterlanServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return errors.NewInterlanServerError(err.Error())
	}
	return nil
}
