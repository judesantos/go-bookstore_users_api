package users

// users.controller.go

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/judesantos/go-bookstore_users_api/services"
	"github.com/judesantos/go-bookstore_users_api/utils/errors"

	"github.com/judesantos/go-bookstore_users_api/domain/users"
)

func getUserId(userId string) (int64, *errors.RestError) {
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, errors.BadRequestError("Invalid user id")
	}

	return id, nil
}

func CreateUser(c *gin.Context) {

	var user users.User
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.BadRequestError("Invalid json body")
		c.JSON(err.Status, err)
		return
	}

	fmt.Println(user)

	result, err := services.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func GetUser(c *gin.Context) {

	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {

	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.BadRequestError("Invalid json body")
		c.JSON(err.Status, err)
		return
	}

	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch

	result, svcErr := services.UsersService.UpdateUser(isPartial, user)
	if svcErr != nil {
		c.JSON(svcErr.Status, err)
		return
	}

	c.JSON(http.StatusAccepted, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func DeleteUser(c *gin.Context) {

	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, "deleted")

}

func Search(c *gin.Context) {

	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {

	var req = users.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := errors.InvalidParameterError("invalid username or password")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UsersService.LoginUser(req)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
