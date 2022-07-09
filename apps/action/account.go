package action

import (
	"strconv"

	. "github.com/LangPham/mila_cast"
	. "github.com/LangPham/mila_go/apps/model"
	. "github.com/LangPham/mila_go/apps/repo"
	"github.com/LangPham/mila_go/util"
	// . "github.com/LangPham/mila_go/util"
	"github.com/gofiber/fiber/v2"
)

func ListAccount() (accounts []Account) {
	Repo.Find(&accounts)
	return
}

func CreateAccount(c *fiber.Ctx) (exchange Exchange) {
	account := new(Account)
	exchange = account.Change(c)

	if exchange.Valid {
		acc := exchange.Data.(Account)
		Repo.Create(&acc)
		exchange.ResultID = strconv.Itoa(int(acc.ID))
	}
	return
}

func GetAccount(id interface{}) (result Account) {
	Repo.First(&result, id)
	return
}

func CheckAccount(c *fiber.Ctx) (result Account) {
	// value := c.Body()
	value := c.FormValue("UserName")
	util.Dump(value)
	// util.Dump(value)
	// Repo.First(&result, 1)
	Repo.Where("user_name = ?", value).Find(&result)
	return
}


func UpdateAccount(c *fiber.Ctx) (exchange Exchange) {
	id := c.Params("id")
	account := GetAccount(id)
	exchange = account.Change(c)
	if exchange.Valid {
		acc := exchange.Data.(Account)
		Repo.Save(&acc)
		exchange.ResultID = strconv.Itoa(int(acc.ID))
	}
	return
}

func DeleteAccount(id interface{}) (result Account) {
	Repo.Delete(&Account{}, id)
	return
}
