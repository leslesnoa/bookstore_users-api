package contorollers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leslesnoa/bookstore_users-api/domain/users"
	"github.com/leslesnoa/bookstore_users-api/services"
	resterr "github.com/leslesnoa/bookstore_users-api/utils/errors"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle Error
		fmt.Println(err.Error())
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: Handle json error
		fmt.Println(err.Error())
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: return bad request to the caller
		fmt.Println(err)
		restErr := resterr.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: handle user creation error
		fmt.Println(err.Error())
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	fmt.Println(string(bytes))
	fmt.Println(err)
	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// func FindUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")

// }
