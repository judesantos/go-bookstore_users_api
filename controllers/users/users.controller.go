package users

// users.controller.go

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/judesantos/go-bookstore_oauth/oauth"
	"github.com/judesantos/go-bookstore_users_api/services"
	"github.com/judesantos/go-bookstore_utils/rest_errors"

	"github.com/judesantos/go-bookstore_users_api/domain/users"
)

func getUserId(userId string) (int64, rest_errors.IRestError) {
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, rest_errors.BadRequestError("Invalid user id")
	}

	return id, nil
}

//
// CreateUser
//
func CreateUser(c *gin.Context) {

	var user users.User
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&user); err != nil {
		err := rest_errors.BadRequestError("Invalid json body")
		c.JSON(err.Status(), err)
		return
	}

	fmt.Println(user)

	result, err := services.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

//
// GetUser
//
func GetUser(c *gin.Context) {

	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	if oauth.GetUserId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshall(false))
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

//
// UpdateUser
//
func UpdateUser(c *gin.Context) {

	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := rest_errors.BadRequestError("Invalid json body")
		c.JSON(err.Status(), err)
		return
	}

	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch

	result, svcErr := services.UsersService.UpdateUser(isPartial, user)
	if svcErr != nil {
		c.JSON(svcErr.Status(), err)
		return
	}

	c.JSON(http.StatusAccepted, result.Marshall(c.GetHeader("X-Public") == "true"))

}

//
// DeleteUser
//
func DeleteUser(c *gin.Context) {

	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, "deleted")

}

//
// Search
//
func Search(c *gin.Context) {

	status := c.Query("status")

	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

//
// Login
//
func Login(c *gin.Context) {

	var req = users.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		restErr := rest_errors.InvalidParameterError("invalid username or password")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user, err := services.UsersService.LoginUser(req)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}
