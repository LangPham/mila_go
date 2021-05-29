package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	MixinThing
	SubjectOf []Article `gorm:"many2many:tags_articles;"`
	TagType   string
	ParentID  *uint
	Child     []Tag `gorm:"foreignkey:ParentID"`
}

//
//func (models Tag) Change(attr []byte) (exchange Exchange) {
//	exchange = Cast(models, attr, "Name", "Url", "Image", "Description", "TagType")
//	exchange.ValidateRequired("Name", "Url")
//	return
//}
