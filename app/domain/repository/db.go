package repository

import (
	"fmt"
	"stncCms/app/domain/entity"
	"stncCms/app/services"

	"github.com/hypnoglow/gormzap"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	_ "gorm.io/driver/mysql"
)

var DB *gorm.DB

//Repositories strcut
type Repositories struct {
	User           services.UserAppInterface
	Post           services.PostAppInterface
	Cat            services.CatAppInterface
	CatPost        services.CatPostAppInterface
	Lang           services.LanguageAppInterface
	WebArchive     services.WebArchiveAppInterface
	WebArchiveLink services.WebArchiveLinksAppInterface

	DB *gorm.DB
}

//DbConnect initial
/*TODO: burada db verisi pointer olarak işaretlenecek oyle gidecek veri*/
func DbConnect(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName, debug, gormAdvancedLogger string) *gorm.DB {

	//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword) //bu postresql

	//DBURL := "root:sel123C#@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	//DBURL := "root:mysql@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//        root:mysql@/tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local
	DBURL := DbUser + ":" + DbPassword + "@tcp(" + DbHost + ":" + DbPort + ")/" + DbName + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"

	db, err := gorm.Open(Dbdriver, DBURL)

	//nunlar gorm 2 versionunda prfexi falan var
	// db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix:   "krbn_", // table name prefix, table for `User` would be `t_users`
	// 		SingularTable: true,    // use singular table name, table for `User` would be `user` with this option enabled
	// 	},
	// 	// Logger: gorm_logrus.New(),
	// })

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	if debug == "DEBUG" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
		log := zap.NewExample()
		db.SetLogger(gormzap.New(log, gormzap.WithLevel(zap.DebugLevel)))
	} else if debug == "DEBUG" || debug == "TEST" && gormAdvancedLogger == "ENABLE" {
		db.LogMode(true)
	} else if debug == "RELEASE" {
		fmt.Println(debug)
		db.LogMode(false)
	}
	DB = db
	return db
}

//RepositoriesInit initial
func RepositoriesInit(db *gorm.DB) (*Repositories, error) {

	return &Repositories{
		User:           UserRepositoryInit(db),
		Post:           PostRepositoryInit(db),
		Cat:            CatRepositoryInit(db),
		CatPost:        CatPostRepositoryInit(db),
		Lang:           LanguageRepositoryInit(db),
		WebArchive:     WebArchiveRepositoryInit(db),
		WebArchiveLink: WebArchiveLinksRepositoryInit(db),
		DB:             db,
	}, nil
}

//Close closes the  database connection
// func (s *Repositories) Close() error {
// 	return s.db.Close()
// }

//Automigrate This migrate all tables
func (s *Repositories) Automigrate() error {
	s.DB.AutoMigrate(&entity.WebArchive{}, &entity.WebArchiveLinks{}, &entity.User{}, &entity.Post{}, &entity.Categories{}, &entity.CategoryPosts{},
		&entity.Languages{}, &entity.Modules{}, &entity.Notes{},
		&entity.Options{})

	return s.DB.Model(&entity.WebArchiveLinks{}).AddForeignKey("web_archive_id", "web_archives(id)", "CASCADE", "CASCADE").Error // one to many (one=web_archives) (many=WebArchiveLinks)

}
