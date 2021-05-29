package controllers

import (
	"github.com/LangPham/mila_go/apps/repo"
	"github.com/LangPham/mila_go/apps_web/views"
	"github.com/gofiber/fiber/v2"
)

func NewTag(c *fiber.Ctx) error {
	commute := repo.NewExchange("account")
	return views.RenderHTML(c, "account/new", views.M{
		"commute": commute,
	})
}

//func CreateTag(c *fiber.Ctx) error {
//	switch ex := action.CreateTag(c); ex.Valid {
//	case true:
//		return c.Redirect("/admin/account/"+ ex.ResultID)
//	default:
//		return views.RenderHTML(c, "account/new", views.M{
//			"exchange": ex,
//		})
//	}
//
//}

//func ShowTag(c *fiber.Ctx) error {
//	id, _ := c.ParamsInt("id")
//	u , _ := action.GetTag(id)
//	return views.RenderHTML(c, "account/show", views.M{
//		"account": u,
//	})
//}



