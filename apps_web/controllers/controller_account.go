package controllers

import (
	"github.com/LangPham/mila/apps/action"
	"github.com/LangPham/mila/apps/aon"
	. "github.com/LangPham/mila_cast"
	//"github.com/LangPham/mila/apps/repo"
	"github.com/LangPham/mila/apps_web/views"
	"github.com/gofiber/fiber/v2"
)

func NewAccount(c *fiber.Ctx) error {
	exchange := NewExchange("account")
	return views.RenderHTML(c, "account/new", views.M{
		"exchange": exchange,
	})
}

func EditAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	account := action.GetAccount(id)
	exchange := NewExchange(account)
	return views.RenderHTML(c, "account/new", views.M{
		"exchange": exchange,
		"id": id,
	})
}

func CreateAccount(c *fiber.Ctx) error {
	switch ex := action.CreateAccount(c); ex.Valid {
	case true:
		return c.Redirect("/admin/account/" + ex.ResultID)
	default:
		return views.RenderHTML(c, "account/new", views.M{
			"exchange": ex,
		})
	}
	return nil
}

func ShowAccount(c *fiber.Ctx) error {
	id :=  c.Params("id")
	account := action.GetAccount(id)
	return views.RenderHTML(c, "account/show", views.M{
		"account": account,
	})
}

func UpdateAccount(c *fiber.Ctx) error {
	switch ex := action.UpdateAccount(c); ex.Valid {
	case true:
		aon.Dump(ex, "UPDATE TRUE")
		return c.Redirect("/admin/account/" + ex.ResultID)
	default:
		aon.Dump(ex, "UPDATE FALSE")
		return views.RenderHTML(c, "account/new", views.M{
			"exchange": ex,
		})
	}
}

func Delete(c *fiber.Ctx) {
	//id, _ := c.ParamsInt("id")
	id := c.Params("id")
	//account := action.GetAccount(id)
	//if account.ID != 0 {
		action.DeleteAccount(id)
	//}
	_ = c.Redirect("/admin/account/", 200)
}