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

//WebArchiveLinks strcut
type WebArchiveLinks struct {
	ID           uint64     `gorm:"primary_key;auto_increment" json:"id"`
	WebArchiveID uint64     `gorm:"not null;" json:"WebArchiveID"`
	Link         string     `gorm:"size:255;not null;" json:"title" validate:"required"`
	Pdf          string     `gorm:"size:255;not null;" json:"content"`
	Png          string     `gorm:"size:255;not null;" json:"short_content"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

//BeforeSave init
func (f *WebArchiveLinks) BeforeSave() {
	f.Link = html.EscapeString(strings.TrimSpace(f.Link))

}

//Prepare init
func (f *WebArchiveLinks) Prepare() {
	f.Link = html.EscapeString(strings.TrimSpace(f.Link))
	f.CreatedAt = time.Now()
	f.UpdatedAt = time.Now()
}

//Validate fluent validation
func (f *WebArchiveLinks) Validate() map[string]string {
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
