package controllers

import (
	"github.com/LangPham/mila/apps/action"
	"github.com/LangPham/mila/apps_web/views"
	. "github.com/LangPham/mila_cast"
	"github.com/gofiber/fiber/v2"
)

func NewTag(c *fiber.Ctx) error {
	exchange := NewExchange("tag")
	return views.RenderHTML(c, "tag/new", views.M{
		"exchange": exchange,
	})
}


func CreateTag(c *fiber.Ctx) error {
	switch ex := action.CreateTag(c); ex.Valid {
	case true:
		return c.Redirect("/admin/tag/" + ex.ResultID)
	default:
		return views.RenderHTML(c, "tag/new", views.M{
			"exchange": ex,
		})
	}
	return nil
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
