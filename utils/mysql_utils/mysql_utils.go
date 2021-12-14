package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	resterr "github.com/leslesnoa/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *resterr.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return resterr.NewNotFoundError("no record matching given id")
		}
		return resterr.NewInternalServerError("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return resterr.NewBadRequestError("invalid data")
	}
	return resterr.NewInternalServerError("error processing request")
}
