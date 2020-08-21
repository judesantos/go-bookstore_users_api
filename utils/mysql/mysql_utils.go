package mysql_utils

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/judesantos/go-bookstore_utils/logger"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) rest_errors.IRestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return rest_errors.NotFoundError("not found")
		}
		logger.Error("sql query failed", err)
		return rest_errors.InternalServerError(
			"request failed",
			errors.New(err.Error()))
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.BadRequestError("invalid parameter")
	}

	logger.Error("sql error", err)
	return rest_errors.InternalServerError(
		"error processing request",
		errors.New(err.Error()))

}
