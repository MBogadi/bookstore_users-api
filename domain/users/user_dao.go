package users

import (
	"../../datasources/mysql/users_db"
	"../../utils/date"
	"../../utils/errors"
	"../../utils/mysql_utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	preparedInsertUser   = "INSERT INTO users_db.users (FirstName, LastName, EmailID, DateCreated, Status, Password) VALUES(?,?,?,?,?,?)"
	preparedGetUser      = "SELECT FirstName, LastName, EmailID, DateCreated, Status, Password FROM users_db.users WHERE ID = ?"
	preparedUpdateUser   = "UPDATE users_db.users SET FirstName = ?, LastName = ?, EmailID = ?, Status = ?, Password = ? WHERE ID = ?"
	preparedDeleteUser   = "DELETE FROM users_db.users WHERE ID = ?"
	preparedFindByStatus = "SELECT ID, FirstName, LastName, EmailID, DateCreated, Status, Password FROM users_db.users WHERE Status = ?"
)

var (
	userDB = make(map[int64]*User)
)

func init() {
	users_db.OpenDBConn()
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare(preparedInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	user.Status = StatusActive
	stmtResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	uid, err := stmtResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = uid
	// no error
	return nil
}

func (user *User) Get() *errors.RestError {

	stmt, err := users_db.ClientDB.Prepare(preparedGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	stmtResult := stmt.QueryRow(user.Id)

	if err := stmtResult.Scan(&user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Update() *errors.RestError {

	stmt, err := users_db.ClientDB.Prepare(preparedUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	stmtResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	rowsChanged, err := stmtResult.RowsAffected()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	if rowsChanged == 0 {
		log.Println("Update to user affected no rows. user_id:", user.Id)
	}

	if rowsChanged > 1 {
		log.Println("Update to user affected more that 1 row. user_id:", user.Id)
		// consider making it part of error type in utils
		return errors.NewInternalServerError(fmt.Sprintf("Update to user affected more that 1 row. user_id: %v", user.Id))
	}

	return nil
}

func (user *User) Delete() *errors.RestError {

	stmt, err := users_db.ClientDB.Prepare(preparedDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	stmtResult, err := stmt.Exec(user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	rowsChanged, err := stmtResult.RowsAffected()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	if rowsChanged == 0 {
		log.Println("No rows deleted for user_id:", user.Id)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestError) {

	stmt, err := users_db.ClientDB.Prepare(preparedFindByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	defer rows.Close()
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}

	var resultRows = make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		resultRows = append(resultRows, user)
	}

	if len(resultRows) == 0 {
		return nil, errors.NewInternalServerError(fmt.Sprintf("No users in '%s' status", status))
	}

	return resultRows, nil
}
