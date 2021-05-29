package controllers

import (
	"github.com/LangPham/mila_go/apps_web/views"
	"github.com/gofiber/fiber/v2"
)

func NewArticle(c *fiber.Ctx) error {
	return views.RenderHTML(c, "article/new", views.M{
		"articles": "text",
	})
}

//
//func CreateUser(c *fiber.Ctx) error {
//	ctx := context.Background()
//	u, e := action.CreateUser(ctx)
//	debug.Dump(u, "USER")
//	debug.Dump(e, "ERR")
//	c.Send([]byte("test"))
//	return nil
//}