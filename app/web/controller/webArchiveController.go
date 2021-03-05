package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"stncCms/app/domain/entity"
	"stncCms/app/domain/helpers/stncsession"
	"stncCms/app/services"

	"strconv"

	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

//WebArchive constructor
type WebArchive struct {
	webArchiveApp     services.WebArchiveAppInterface
	webArchiveLinkApp services.WebArchiveLinksAppInterface

	userApp services.UserAppInterface
}

const webArchivePath = "admin/webarchive/"

//InitWebArchive webarchive controller constructor
func InitWebArchive(wApp services.WebArchiveAppInterface, wLinkApp services.WebArchiveLinksAppInterface, uApp services.UserAppInterface) *WebArchive {
	return &WebArchive{
		webArchiveApp:     wApp,
		webArchiveLinkApp: wLinkApp,

		userApp: uApp,
	}
}

//Index list
func (access *WebArchive) Index(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	paginator = &pagination.Paginator{}
	var total int64
	access.webArchiveApp.Count(&total)
	PerPage := 8
	paginator := pagination.NewPaginator(c.Request, PerPage, total)
	offset := paginator.Offset()
	posts, _ := access.webArchiveApp.GetAllP(PerPage, offset)
	viewData := pongo2.Context{
		"paginator": paginator,
		"title":     "İçerik Ekleme",
		"posts":     posts,
		"csrf":      csrf.GetToken(c),
	}
	c.HTML(
		http.StatusOK,
		webArchivePath+"index.html",
		viewData,
	)
}

// Create webarchive
// Create webarchive
func (access *WebArchive) Create(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	viewData := pongo2.Context{
		"title": "WebArchive Add",
		"csrf":  csrf.GetToken(c),
	}

	c.HTML(
		http.StatusOK,
		webArchivePath+"create.html",
		viewData,
	)
}

//Store save method
func (access *WebArchive) Store(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var webarchive, _, _ = webArchiveModel(c)
	fmt.Println(webarchive)
	var saveError = make(map[string]string)
	saveError = webarchive.Validate()

	if len(saveError) == 0 {
		saveData, saveErr := access.webArchiveApp.Save(&webarchive)
		if saveErr != nil {
			saveError = saveErr
		}
		lastID := strconv.FormatUint(uint64(saveData.ID), 10)
		linksAll := c.PostForm("LinksAll")
		var re = regexp.MustCompile(`(?m)(?P<origin>(?P<protocol>http[s]?:)?\/\/(?P<host>[a-z0-9A-Z-_.]+))(?P<port>:\d+)?(?P<path>[\/a-zA-Z0-9-\.]+)?(?P<search>\?[^#\n]+)?(?P<hash>#.*)?`)
		for i, match := range re.FindAllString(linksAll, -1) {
			fmt.Println(match, "found at index", i)
			// webarchivelink entity.WebArchiveLinks{}
			var webarchivelink = entity.WebArchiveLinks{}
			webarchivelink.Link = match
			webarchivelink.WebArchiveID = saveData.ID
			access.webArchiveLinkApp.Save(&webarchivelink)
		}

		stncsession.SetFlashMessage("Record successfully added", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/webarchive/edit/"+lastID)
		return
	}

	viewData := pongo2.Context{
		"title":      "webarchive create",
		"csrf":       csrf.GetToken(c),
		"err":        saveError,
		"post":       webarchive,
		"webarchive": webarchive,
	}
	c.HTML(
		http.StatusOK,
		webArchivePath+"create.html",
		viewData,
	)
}

//Edit edit data
func (access *WebArchive) Edit(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	flashMsg := stncsession.GetFlashMessage(c)
	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {
		if posts, err := access.webArchiveApp.GetByID(ID); err == nil {
			viewData := pongo2.Context{
				"title":    "webarchive edit",
				"post":     posts,
				"csrf":     csrf.GetToken(c),
				"flashMsg": flashMsg,
			}
			c.HTML(
				http.StatusOK,
				webArchivePath+"edit.html",
				viewData,
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

//Update data
func (access *WebArchive) Update(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)
	var webarchive, _, id = webArchiveModel(c)

	var savePostError = make(map[string]string)
	savePostError = webarchive.Validate()
	if len(savePostError) == 0 {
		_, saveErr := access.webArchiveApp.Update(&webarchive)
		if saveErr != nil {
			savePostError = saveErr
		}
		stncsession.SetFlashMessage("Record successfully edit", c)
		c.Redirect(http.StatusMovedPermanently, "/admin/webarchive/edit/"+id)
		return
	}

	viewData := pongo2.Context{
		"title": "içerik ekleme",
		"err":   savePostError,
		"csrf":  csrf.GetToken(c),
		"post":  webarchive,
	}
	c.HTML(
		http.StatusOK,
		webArchivePath+"edit.html",
		viewData,
	)
}

//Delete data
func (access *WebArchive) Delete(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {

		access.webArchiveApp.Delete(ID)
		stncsession.SetFlashMessage("Success Delete", c)

		c.Redirect(http.StatusTemporaryRedirect, "/admin/webarchive")
		return

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (access *WebArchive) LinkPdfRun(c *gin.Context) {
	stncsession.IsLoggedInRedirect(c)

	if ID, err := strconv.ParseUint(c.Param("ID"), 10, 64); err == nil {

		access.webArchiveApp.Delete(ID)
		stncsession.SetFlashMessage("Success RUN", c)

		c.Redirect(http.StatusMovedPermanently, "/admin/webarchive/edit/"+c.Param("ID"))
		return

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

//form webarchive model
func webArchiveModel(c *gin.Context) (webarchive entity.WebArchive, idD uint64, idStr string) {
	id := c.PostForm("ID")
	title := c.PostForm("Title")
	linksAll := c.PostForm("LinksAll")
	excerpt := c.PostForm("Excerpt")
	idInt, _ := strconv.Atoi(id)
	var idN uint64
	idN = uint64(idInt)
	//	var webarchive = entity.WebArchive{}
	webarchive.ID = idN
	webarchive.UserID = 1
	webarchive.Title = title
	webarchive.Status = 1
	webarchive.LinksAll = linksAll
	webarchive.Excerpt = excerpt
	return webarchive, idN, id
}
