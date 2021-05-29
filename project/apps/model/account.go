package model

import (
	. "github.com/LangPham/mila/apps/repo"
	"github.com/emirpasic/gods/maps/hashmap"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserName string
	Role     string
	Email	string
}


func (models Account) Change(attr *hashmap.Map) (exchange Exchange) {

	exchange = Cast(models, attr, "UserName", "Role", "Email")
	models.Role = "GUEST"

	exchange.ValidateRequired("UserName")
	return
}
