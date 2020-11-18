package mysql_utils

import (
	"../../utils/errors"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in resultset"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record matching given User ID")
		}
		return errors.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewInternalServerError("Duplicate record")
	}
	return errors.NewInternalServerError("error processing request")

}
