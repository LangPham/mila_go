package action

import (
	. "github.com/LangPham/mila/apps/model"
	. "github.com/LangPham/mila/apps/repo"
	. "github.com/LangPham/mila_cast"
	"github.com/gofiber/fiber/v2"
	"strconv"
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
