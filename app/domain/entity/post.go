package entity

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	tr_translations "gopkg.in/go-playground/validator.v9/translations/tr"
)


//Post strcut
type Post struct {
	ID           uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint64     `gorm:"not null;" json:"user_id"`
	PostTitle    string     `gorm:"size:255;not null;" json:"title" validate:"required"`
	PostContent  string     `gorm:"type:text;" json:"content" validate:"required"`
	PostExcerpt  string     `gorm:"type:text;" json:"short_content"`
	PostPassword string     `gorm:"size:255;null;" json:"password"`
	PostSlug     string     `gorm:"size:255;null;" json:"slug"`
	MenuOrder    string     `gorm:"size:255;null;" json:"order"`
	CommentCount string     `gorm:"size:255;null;" json:"cmment_count"`
	PostType     int        `gorm:"type:smallint unsigned;NOT NULL;DEFAULT:'1'" validate:"required"`
	PostStatus   int        `gorm:"type:smallint unsigned;NOT NULL;DEFAULT:'1'" validate:"required"`
	Picture      string     `gorm:"size:255;null;" json:"picture" `
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

//BeforeSave init
func (f *Post) BeforeSave() {
	f.PostTitle = html.EscapeString(strings.TrimSpace(f.PostTitle))
	f.PostExcerpt = html.EscapeString(strings.TrimSpace(f.PostExcerpt))
	f.PostSlug = html.EscapeString(strings.TrimSpace(f.PostSlug))
}

/*
func (post *Post) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
 }
*/

//Prepare init
func (f *Post) Prepare() {
	f.PostTitle = html.EscapeString(strings.TrimSpace(f.PostTitle))
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

//Validate fluent validation
func (f *Post) Validate() map[string]string {
	var (
		validate *validator.Validate
		uni      *ut.UniversalTranslator
	)
	tr := en.New()
	uni = ut.New(tr, tr)
	trans, _ := uni.GetTranslator("tr")
	validate = validator.New()
	tr_translations.RegisterDefaultTranslations(validate, trans)

	errorLog := make(map[string]string)

	err := validate.Struct(f)
	fmt.Println(err)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		fmt.Println(errs)
		for _, e := range errs {
			// can translate each error one at a time.
			lng := strings.Replace(e.Translate(trans), e.Field(), "Burası", 1)
			errorLog[e.Field()+"_error"] = e.Translate(trans)
			// errorLog[e.Field()] = e.Translate(trans)
			errorLog[e.Field()] = lng
			errorLog[e.Field()+"_valid"] = "is-invalid"
		}
	}
	return errorLog
}
