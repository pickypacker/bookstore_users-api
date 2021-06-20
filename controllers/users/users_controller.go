package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pickypacker/bookstore_users-api/domain/users"
)

func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "Not implemented")
}

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)

	if err == nil {

		err = json.Unmarshal(bytes, &user)
		if err == nil {

		} else {
			fmt.Println(err)
		}
	}
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}
