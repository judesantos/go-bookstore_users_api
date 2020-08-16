package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/judesantos/go-bookstore_users_api/logger"
	"github.com/judesantos/go-bookstore_users_api/utils/errors"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NotFoundError("not found")
		}
		logger.Error("sql query failed", err)
		return errors.InternalServerError("request failed")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.BadRequestError("invalid parameter")
	}

	logger.Error("sql error", err)
	return errors.InternalServerError("error processing request")
}
