package model

import (
	. "github.com/LangPham/mila/apps/repo"
	"github.com/gofiber/fiber/v2"
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

func (models Tag) Change(c *fiber.Ctx) (exchange Exchange) {
	exchange = Cast(models, c)
	exchange.ValidateModel()
	return
}
