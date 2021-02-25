package dto

//KurbanTableName table name
var KurbanTableName string = "kurbanlar"

//Kurban dto
type Kurban struct {
	ID      uint64
	AdSoyad string
}

/*
//GenelKurbanRefDataDTO genel kurban ref
func GenelKurbanRefDataDTO(data *GenelKurban) *entity.GenelKurban {
	return &entity.GenelKurban{
		ID:      data.ID,
		AdSoyad: data.AdSoyad,
	}
}
*/
