package model

import (
	. "github.com/LangPham/mila_cast"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserName string `cast:"user_name" validate:"required"`
	Email    string `cast:"email" validate:"required,email"`
	Role     string `validate:"required"`
}

func (models Account) Change(c *fiber.Ctx) (exchange Exchange) {

	exchange = Cast(models, c)
	exchange.PutField("Role", "GUEST")
	exchange.ValidateModel()

	return
}
