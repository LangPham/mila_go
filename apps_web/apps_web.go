package apps_web

import (
	"github.com/LangPham/mila/apps/aon"
	"github.com/LangPham/mila/apps_web/controllers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func AppsWeb() {
	app := fiber.New()
	app.Static("/", "./public/static")
	//app.Get("/", controllers.PageIndex)
	app.Use(func(c *fiber.Ctx) error {

		switch c.Method() {
		case "POST":
			if c.FormValue("_METHOD") != "" {
				c.Method(c.FormValue("_METHOD"))
			}
		default:
			aon.Dump(c.Method() , "MID")
		}
		// Set some security headers:
		//c.Set("X-XSS-Protection", "1; mode=block")
		//c.Set("X-Content-Type-Options", "nosniff")
		//c.Set("X-Download-Options", "noopen")
		//c.Set("Strict-Transport-Security", "max-age=5184000")
		//c.Set("X-Frame-Options", "SAMEORIGIN")
		//c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})

	admin := app.Group("/admin")
	// account
	account := admin.Group("/account")
	account.Get("/new", controllers.NewAccount)
	account.Get("/:id/edit", controllers.EditAccount)
	account.Get("/:id", controllers.ShowAccount)
	account.Post("/", controllers.CreateAccount)
	account.Put("/:id", controllers.UpdateAccount)

	tag:= admin.Group("tag")
	tag.Get("/new", controllers.NewTag)
	tag.Post("/", controllers.CreateTag)
	//admin.Get("/user/:id", controllers.ShowUser)
	// tag
	//admin.Get("/account/new", controllers.NewAccount)
	//admin.Post("/account", controllers.CreateAccount)

	log.Fatal(app.Listen(":8080"))
}
