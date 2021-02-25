package stncsession

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// IsLoggedIn checks if the current user is logged in
func IsLoggedIn(c *gin.Context) bool {
	session := sessions.Default(c)
	return session.Get("userID") != nil
}

//IsLoggedInRedirect redirect site
func IsLoggedInRedirect(c *gin.Context) {
	l := IsLoggedIn(c)
	if l == false {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}

//GetUserID get userID data
func GetUserID(c *gin.Context) (userID string) {
	_userID := GetSession("userID", c)
	if _userID == "" {
		userID = "null"
	} else {
		userID = _userID
		fmt.Println("data " + userID)
	}
	return userID
}

// SetStoreUserID stores the userId for teh current user
func SetStoreUserID(c *gin.Context, userID uint64) {
	session := sessions.Default(c)
	session.Set("userID", userID)
	session.Save()
}

// ClearUserID clears the userId for the current suer //ClearUserIDFromCookie
func ClearUserID(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("userID", nil)
	session.Save()
}
