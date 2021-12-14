package contorollers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leslesnoa/bookstore_users-api/domain/users"
	"github.com/leslesnoa/bookstore_users-api/services"
	resterr "github.com/leslesnoa/bookstore_users-api/utils/errors"
)

var (
	counter int
)

func getUserId(userIdParam string) (int64, *resterr.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, resterr.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println("start CreateUserFunc.")
	// fmt.Println(user)
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO: Handle Error
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO: Handle json error
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(user)
	// fmt.Println(string(bytes))
	// if err := c.ShouldBindJSON(&user); err != nil {
	// 	//TODO: return bad request to the caller
	// 	fmt.Println(err.Error())
	// 	restErr := resterr.NewBadRequestError("invalid json body")
	// 	c.JSON(restErr.Status, restErr)
	// 	return
	// }

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: handle user creation error
		fmt.Println(saveErr)
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// func FindUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")

// }

func UpdateUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := resterr.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
