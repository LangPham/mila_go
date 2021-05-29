package controllers

import (
	"github.com/LangPham/mila/apps/action"
	"github.com/LangPham/mila/apps/repo"
	"github.com/LangPham/mila/apps_web/views"
	debug "github.com/LangPham/mila_debug"
	"github.com/gofiber/fiber/v2"
)

func NewAccount(c *fiber.Ctx) error {
	exchange := repo.NewExchange("account")

	return views.RenderHTML(c, "account/new", views.M{
		"exchange": exchange,
	})
}

func EditAccount(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	account := action.GetAccount(id)
	exchange := repo.NewExchange(account)
	return views.RenderHTML(c, "account/new", views.M{
		"exchange": exchange,
	})
}

func CreateAccount(c *fiber.Ctx) error {
	switch ex := action.CreateAccount(c); ex.Valid {
	case true:
		return c.Redirect("/admin/account/"+ ex.ResultID)
	default:
		return views.RenderHTML(c, "account/new", views.M{
			"exchange": ex,
		})
	}
	return nil
}

func ShowAccount(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	account := action.GetAccount(id)
	debug.Dump(account, "AAC")
	return views.RenderHTML(c, "account/show", views.M{
		"account": account,
	})
}



