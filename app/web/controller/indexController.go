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

	option4 := entity.Options{OptionName: "satis_birim_fiyati_1", OptionValue: ""}
	db.Debug().Create(&option4)

	option5 := entity.Options{OptionName: "satis_birim_fiyati_2", OptionValue: ""}
	db.Debug().Create(&option5)

	option6 := entity.Options{OptionName: "satis_birim_fiyati_3", OptionValue: ""}
	db.Debug().Create(&option6)

	option7 := entity.Options{OptionName: "dusuk_agirlik_kilo", OptionValue: ""}
	db.Debug().Create(&option7)

	option78 := entity.Options{OptionName: "orta_agirlik_kilo", OptionValue: ""}
	db.Debug().Create(&option78)

	option786 := entity.Options{OptionName: "yuksek_agirlik_kilo", OptionValue: ""}
	db.Debug().Create(&option786)

	option8 := entity.Options{OptionName: "alis_birim_fiyati_1", OptionValue: ""}
	db.Debug().Create(&option8)

	option9 := entity.Options{OptionName: "alis_birim_fiyati_2", OptionValue: ""}
	db.Debug().Create(&option9)

	option10 := entity.Options{OptionName: "alis_birim_fiyati_3", OptionValue: ""}
	db.Debug().Create(&option10)

	option11 := entity.Options{OptionName: "otomatik_sira_2021", OptionValue: "1"}
	db.Debug().Create(&option11)

	user := entity.User{FirstName: "Sel", LastName: "t", Email: "selmantunc@gmail.com", Password: "$2a$10$QPiWAgMpwHBkDjBL5pPd2.HBlfdniuGOvZd5kh.ILLjKFo67rvfsO"}
	db.Debug().Create(&user)

	c.JSON(http.StatusOK, "yapıldı")
}
