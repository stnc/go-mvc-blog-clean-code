//  golang gin framework mvc and clean code project
//  Licensed under the Apache License 2.0
//  @author Selman TUNÇ <selmantunc@gmail.com>
//  @link: https://github.com/stnc/go-mvc-blog-clean-code
//  @license: Apache License 2.0
package main

import (
	"log"
	"net/http"
	"os"

	"stncCms/app/domain/repository"
	"stncCms/app/infrastructure/auth"

	apiController "stncCms/app/web.api/controller"
	"stncCms/app/web.api/controller/fileupload"
	"stncCms/app/web.api/controller/middleware"
	"stncCms/app/web/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stnc/pongo2gin"
	csrf "github.com/utrack/gin-csrf"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// err := beeep.Alert("Uygulama çalıştı", "Web Server Run localhost:"+port, "assets/warning.png")
	// if err != nil {
	// 	panic(err)
	// }

	//redis details
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	debugMode := os.Getenv("MODE")
	gormAdvancedLogger := os.Getenv("GORM_ZAP_LOGGER")
	db := repository.DbConnect(dbdriver, user, password, port, host, dbname, debugMode, gormAdvancedLogger)
	services, err := repository.RepositoriesInit(db)
	if err != nil {
		panic(err)
	}
	//defer services.Close()
	services.Automigrate()

	redisService, err := auth.RedisDBInit(redisHost, redisPort, redisPassword)
	if err != nil {
		log.Fatal(err)
	}

	token := auth.NewToken()

	upload := fileupload.NewFileUpload()

	usersAPI := apiController.InitUsers(services.User, redisService.Auth, token)

	postsAPI := apiController.InitPost(services.Post, services.User, redisService.Auth, token)

	posts := controller.InitPost(services.Post, services.CatPost, services.Cat, services.Lang, services.User, upload)

	login := controller.InitLogin(services.User)

	authenticate := apiController.NewAuthenticate(services.User, redisService.Auth, token)

	webArchive := controller.InitWebArchive(services.WebArchive, services.WebArchiveLink, services.User)

	switch debugMode {
	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)

	case "DEBUG":
		gin.SetMode(gin.DebugMode)

	case "TEST":
		gin.SetMode(gin.TestMode)

	default:
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(gin.Recovery())

	//TODO: https://github.com/denisbakhtin/ginblog/blob/master/main.go burada memstore kullanımı var ona bakılablir

	store := cookie.NewStore([]byte("SpeedyGonzales"))

	store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: 7 * 86400}) //Also set Secure: true if using SSL, you should though

	r.Use(sessions.Sessions("myCRM", store))

	r.Use(middleware.CORSMiddleware()) //For CORS

	//TODO: csrf kontrolu nasıl olacak
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "SpeedyGonzales",
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.HTMLRender = pongo2gin.TemplatePath("public/view")

	r.MaxMultipartMemory = 1 >> 20 // 8 MiB

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Static("/assets", "./public/static")

	r.StaticFS("/upload", http.Dir("./public/upl"))
	//r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("/", controller.Index)
	r.GET("login", login.Login)
	r.POST("login", login.LoginPost)
	r.GET("logout", login.Logout)

	r.GET("optionsDefault", controller.OptionsDefault)

	//api routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("users", usersAPI.SaveUser)
		v1.GET("users", usersAPI.GetUsers)

		v1.GET("users/:user_id", usersAPI.GetUser)
		v1.GET("postall", postsAPI.GetAllPost)
		v1.POST("post", postsAPI.SavePost)
		v1.PUT("post/:post_id", middleware.AuthMiddleware(), postsAPI.UpdatePost)
		v1.GET("post/:post_id", postsAPI.GetPostAndCreator)
		v1.DELETE("post/:post_id", middleware.AuthMiddleware(), postsAPI.DeletePost)
		v1.POST("login", authenticate.Login)
		v1.POST("logout", authenticate.Logout)
		v1.POST("refresh", authenticate.Refresh)
	}

	adminPost := r.Group("/admin/post")
	{

		adminPost.GET("/", posts.Index)
		adminPost.GET("index", posts.Index)
		adminPost.GET("create", posts.Create)
		adminPost.POST("store", posts.Store)
		adminPost.GET("edit/:postID", posts.Edit)
		adminPost.POST("update", posts.Update)

	}

	adminWebarchive := r.Group("/admin/webarchive")
	{
		adminWebarchive.GET("/", webArchive.Index)
		adminWebarchive.GET("/index", webArchive.Index)
		adminWebarchive.GET("create", webArchive.Create)
		adminWebarchive.POST("store", webArchive.Store)
		adminWebarchive.GET("edit/:ID", webArchive.Edit)
		adminWebarchive.POST("update", webArchive.Update)
		adminWebarchive.GET("delete/:ID", webArchive.Delete)
		adminWebarchive.GET("run/:ID", webArchive.LinkPdfRun)
	}

	//Starting the application
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8888" //localhost
	}
	log.Fatal(r.Run(":" + appPort))

}
