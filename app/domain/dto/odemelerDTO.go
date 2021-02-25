package dto

import "time"

//OdemelerTableName table name
var OdemelerTableName string = "odemeler"

//Odemeler dto layer
type Odemeler struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	KurbanID uint64 `gorm:"not null;" json:"kurban_id"`
	UserID   uint64 `gorm:"not null;" json:"user_id"`
	Aciklama string `gorm:"type:text;" json:"aciklama"`
	//durum=  taksit eklemişse durum 1 / İlk Eklenen Fiyat değeri 2 / kasa borçlu kalmışsa değer 3 olur
	BorcDurum      int       `gorm:"type:smallint unsigned;NOT NULL;DEFAULT:'1'" validate:"required"`
	VerilenUcret   string    `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';"  validate:"required,numeric" `
	KalanUcret     string    `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';" validate:"numeric" `
	KasaBorcu      float64   `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';" validate:"numeric" `
	VerildigiTarih time.Time `gorm:"type:datetime" json:"verildigiTarih"`
}

//OdemelerSonFiyat dto layer
type OdemelerSonFiyat struct {
	ID           uint64  `gorm:"primary_key;auto_increment" json:"id"`
	VerilenUcret float64 `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';"  validate:"required,numeric" `
	KalanUcret   float64 `gorm:"type:decimal(10,2); NOT NULL; DEFAULT:'0';" validate:"required,numeric" `
}

//TableName override
func (gk *Odemeler) TableName() string {
	return OdemelerTableName
}
