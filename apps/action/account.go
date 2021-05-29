package action

import (
	. "github.com/LangPham/mila_go/apps/model"
	. "github.com/LangPham/mila_go/apps/repo"
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

func GetAccount(id int) (result Account) {
	Repo.First(&result, id)
	return
}

//
//func UpdateAccount(account Account, attrs interface{}) (error bool, result Change) {
//	var changeset Change
//
//	changeset = account.Change(attrs)
//	account = changeset.Data.(Account)
//
//	if changeset.Valid {
//		Repo.Save(&account)
//		changeset.IdResult = strconv.Itoa(int(account.ID))
//
//		result = changeset
//		error = true
//	} else {
//		result = changeset
//		error = false
//	}
//	return
//}
//
//func DeleteAccount(id string) (result Account) {
//	Repo.Delete(&Account{}, id)
//	return
//}
