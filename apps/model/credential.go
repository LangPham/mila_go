package model

import (
	. "github.com/LangPham/mila_cast"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Credential struct {
	gorm.Model
	Genre string `cast:"genre" validate:"required"`
	Email string `cast:"email" validate:"required,email"`
	AccountID uint
}

func (models Credential) Change(c *fiber.Ctx) (exchange Exchange) {

	exchange = Cast(models, c)
	exchange.PutField("Role", "GUEST")
	exchange.ValidateModel()

	return
}
