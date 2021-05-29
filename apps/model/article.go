package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	MixinThing
	Body   string
	Abouts []Tag `gorm:"many2many:tags_articles;"`
}

//func (models Article) Change(attr []byte) (exchange Exchange) {
//	exchange = Cast(models, attr, "Name", "Url" , "Image", "Description", "Body")
//	exchange.ValidateRequired("Name", "Url", "Body")
//	return
//}
