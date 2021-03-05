package controller

import (
	"net/http"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/services"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

//Options constructor
type Options struct {
	OptionsApp services.OptionsAppInterface
}

const viewPathOptions = "admin/options/"

//InitOptions post controller constructor
func InitOptions(OptionsApp services.OptionsAppInterface) *Options {
	return &Options{
		OptionsApp: OptionsApp,
	}
}

//Index list
func (access *Options) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	hisseAdeti := access.OptionsApp.GetOption("hisse_adeti")
	SatisBirimFiyati1 := access.OptionsApp.GetOption("satis_birim_fiyati_1")
	SatisBirimFiyati2 := access.OptionsApp.GetOption("satis_birim_fiyati_2")
	SatisBirimFiyati3 := access.OptionsApp.GetOption("satis_birim_fiyati_3")
	AlisBirimFiyati1 := access.OptionsApp.GetOption("alis_birim_fiyati_1")
	AlisBirimFiyati2 := access.OptionsApp.GetOption("alis_birim_fiyati_2")
	AlisBirimFiyati3 := access.OptionsApp.GetOption("alis_birim_fiyati_3")

	viewData := pongo2.Context{
		"title":                "Ayarlar",
		"csrf":                 csrf.GetToken(c),
		"hisse_adeti":          hisseAdeti,
		"satis_birim_fiyati_1": SatisBirimFiyati1,
		"satis_birim_fiyati_2": SatisBirimFiyati2,
		"satis_birim_fiyati_3": SatisBirimFiyati3,
		"alis_birim_fiyati_1":  AlisBirimFiyati1,
		"alis_birim_fiyati_2":  AlisBirimFiyati2,
		"alis_birim_fiyati_3":  AlisBirimFiyati3,
	}

	c.HTML(
		http.StatusOK,
		viewPathOptions+"index.html",
		viewData,
	)
}

//Update list
func (access *Options) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	access.OptionsApp.SetOption("hisse_adeti", c.PostForm("hisse_adeti"))
	access.OptionsApp.SetOption("satis_birim_fiyati_1", c.PostForm("satis_birim_fiyati_1"))
	access.OptionsApp.SetOption("satis_birim_fiyati_2", c.PostForm("satis_birim_fiyati_2"))
	access.OptionsApp.SetOption("satis_birim_fiyati_3", c.PostForm("satis_birim_fiyati_3"))

	c.Redirect(http.StatusMovedPermanently, "/admin/options")
	return

}
