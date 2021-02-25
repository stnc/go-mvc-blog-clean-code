package stncsession

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// alertType  = success,danger,info,warning

//SetFlashMessage init
func SetFlashMessage(msg string, c *gin.Context) {

	session := sessions.Default(c)
	// session.Set("flashMsgAlertType", alertType)
	// session.Save()
	// session.Set("flashMsg", msg)
	// session.Save()
	session.AddFlash(msg, "alertType")
	session.Save()
}

//GetFlashMessage get flash mesaj
func GetFlashMessage(c *gin.Context) string {

	session := sessions.Default(c)
	// flashMsg := session.Get("flashMsg")
	// flashMsgAlertType := session.Get("flashMsgAlertType")

	// if flashMsg != nil {
	// 	return flashMsg.(string), flashMsgAlertType.(string)
	// } else {
	// 	return "", ""
	// }
	var str string
	str = fmt.Sprintf("%v", session.Flashes("alertType"))
	session.Delete("alertType")
	session.Save()
	if str == "[]" {
		str = ""
	}
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	return str
}

//LoginControl login contr
func LoginControl(c *gin.Context) {
	if GetUserID(c) == "null" {
		c.Redirect(http.StatusTemporaryRedirect, "/admin/login")
	}
}

/*
r.GET("/setFlashMessage", set)
r.GET("/getFlashMessage", get)

func set(c *gin.Context) {
	fm := "ini flash message"
	fm2 := "ini flash message2"
	stncsession.SetFlashMessage(fm, fm2, c)
	// flashMsg, flashMsgAlertType := stncsession.FlashMsgGet(c)
}

func get(c *gin.Context) {
	fm, fm2 := stncsession.GetFlashMessage(c)

	fmt.Println(fm)
	fmt.Println(fm2)
}


*/

/*
//FlashMsgSetter e
func FlashMsgSetter(msg string, alertType string, c echo.Context, time ...int) {

	sess, _ := session.Get("flashMsg", c)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 1,
		//	   MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["flashMsgAlertType"] = alertType
	sess.Values["flashMsg"] = msg
	sess.Save(c.Request(), c.Response())

}

//FlashMsgGet get flash mesaj
func FlashMsgGet(c echo.Context) (string, string) {
	//var flashMsg, flashMsgAlertType string
	sess, _ := session.Get("flashMsg", c)
	flashMsg := sess.Values["flashMsg"]
	flashMsgAlertType := sess.Values["flashMsgAlertType"]

	if flashMsg != nil {
		return flashMsg.(string), flashMsgAlertType.(string)
	} else {
		return "", ""
	}

}


*/

/*
//UserInfoSet User Informaiton
func UserInfoSet(msg string, c echo.Context, time ...int) {

	sess, _ := session.Get("userInfo", c)
	sess.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 1,
		//	   MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["userInfo"] = msg
	sess.Save(c.Request(), c.Response())

}

//UserIDSet User Informaiton
func UserIDSet(id int, c echo.Context, time ...int) {
	sess, _ := session.Get("userID", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["userID"] = id
	sess.Save(c.Request(), c.Response())
}

//UserIDGet User Informaiton
func UserIDGet(c echo.Context) int {
	sess, _ := session.Get("userID", c)
	userID := sess.Values["userID"]
	return userID.(int)
}

// ReDirectSetter redirect
func ReDirectSetter(value string, c echo.Context, time ...int) {
	sess, _ := session.Get("Redirect", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["redirect"] = value
	sess.Save(c.Request(), c.Response())
}

// ReDirectDelete del
func ReDirectDelete(c echo.Context) {
	sess, _ := session.Get("Redirect", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["redirect"] = "/"
	sess.Save(c.Request(), c.Response())
}
*/
