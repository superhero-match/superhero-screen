package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckEmail checks if email already registered.
// This check is done after the user has logged-in.
// If email is registered, return status OK and true for registered.
// If not, return status OK and false for registered, and user will be redirected to register screen.
func (ctl *Controller) CheckEmail(c *gin.Context) {
	var isRegistered bool
	var isDeleted bool
	var isBlocked bool

	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":       http.StatusInternalServerError,
			"isRegistered": isRegistered,
			"isDeleted":    isDeleted,
			"isBlocked":    isBlocked,
		})

		return
	}

	//err := ctl.ES.DeleteIndex()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//err = ctl.ES.CreateIndex()
	//if err != nil {
	//	fmt.Println(err)
	//}

	resp, err := ctl.ES.CheckEmailExists(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":       http.StatusInternalServerError,
			"isRegistered": isRegistered,
			"isDeleted":    isDeleted,
			"isBlocked":    isBlocked,
		})

		return
	}

	fmt.Println()
	fmt.Printf("resp: %+v", resp)
	fmt.Println()

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"isRegistered": resp.IsRegistered,
		"isDeleted":    resp.IsDeleted,
		"isBlocked":    resp.IsBlocked,
	})
}
