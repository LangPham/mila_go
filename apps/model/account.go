package model

import (
	. "github.com/LangPham/mila_go/apps/repo"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserName string `cast:"user_name" validate:"required"`
	Email string `cast:"email" validate:"required,email"`
	Role     string `validate:"required"`
	//Email    string `cast:"email" validate:"required,email"`
	//Age int `cast:"age" validate:"gt=10"`
}

func (models Account) Change(c *fiber.Ctx) (exchange Exchange) {

	exchange = Cast(models, c)
	exchange.PutField("Role", "GUEST")
	exchange.ValidateModel()

	return
}
