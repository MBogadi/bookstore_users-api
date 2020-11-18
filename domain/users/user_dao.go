package users

import (
	"../../datasources/mysql/users_db"
	"../../utils/date"
	"../../utils/errors"
	"../../utils/mysql_utils"
	_ "github.com/go-sql-driver/mysql"
)

const (
	preparedInsertUser = "INSERT INTO users_db.users (FirstName, LastName, EmailID, DateCreated) VALUES(?,?,?,?)"
	preparedGetUser    = "SELECT FirstName, LastName, EmailID, DateCreated FROM users_db.users WHERE ID=?"
)

var (
	userDB = make(map[int64]*User)
)

func init() {
	users_db.OpenDBConn()
}

func (user *User) Get() *errors.RestError {

	stmt, err := users_db.ClientDB.Prepare(preparedGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	stmtResult := stmt.QueryRow(user.Id)

	if err := stmtResult.Scan(&user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
		//return errors.NewInternalServerError(fmt.Sprintf("Failed to get user record for UserID %d. Error: %s", user.Id, err.Error()))
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.ClientDB.Prepare(preparedInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	stmtResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
		//return errors.NewInternalServerError(err.Error())
	}

	uid, err := stmtResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
		//return errors.NewInternalServerError(err.Error())
	}
	user.Id = uid
	// no error
	return nil
}
