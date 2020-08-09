package users

import (
	"github.com/gin-gonic/gin"
	"github.com/judesantos/go-bookstore_users_api/services"
	"github.com/judesantos/go-bookstore_users_api/utils/errors"
	"net/http"
	"strconv"

	"github.com/judesantos/go-bookstore_users_api/domain/users"
)

func CreateUser(c *gin.Context) {

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.BadRequestError("Invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.BadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}