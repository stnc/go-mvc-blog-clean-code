package controller

import (
	"net/http"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/domain/repository"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

const viewPathIndex = "admin/index/"

//Index all list f
func Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	viewData := pongo2.Context{
		"title": "Posts",
		"csrf":  csrf.GetToken(c),
	}
	c.HTML(
		http.StatusOK,
		viewPathIndex+"index.html",
		viewData,
	)
}

//OptionsDefault all list f
func OptionsDefault(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	db := repository.DB

	option1 := entity.Options{OptionName: "siteurl", OptionValue: ""}
	db.Debug().Create(&option1)

	option2 := entity.Options{OptionName: "kurban_yili", OptionValue: "2021"}
	db.Debug().Create(&option2)

	option3 := entity.Options{OptionName: "hisse_adeti", OptionValue: "7"}
	db.Debug().Create(&option3)

	cat := entity.Categories{Name: "News"}
	db.Debug().Create(&cat)

	user := entity.User{FirstName: "Sel", LastName: "t", Email: "selmantunc@gmail.com", Password: "$2a$10$QPiWAgMpwHBkDjBL5pPd2.HBlfdniuGOvZd5kh.ILLjKFo67rvfsO"}
	db.Debug().Create(&user)

	c.JSON(http.StatusOK, "done")
}
