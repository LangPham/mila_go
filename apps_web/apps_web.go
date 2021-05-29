package apps_web

import (
	"github.com/LangPham/mila_go/apps_web/controllers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func AppsWeb() {
	app := fiber.New()
	app.Static("/", "./public/static")
	//app.Get("/", controllers.PageIndex)

	admin := app.Group("/admin")
	// account
	account := admin.Group("/account")
	account.Get("/new", controllers.NewAccount)
	account.Get("/:id/edit", controllers.EditAccount)
	account.Get("/:id", controllers.ShowAccount)
	account.Post("/", controllers.CreateAccount)
	//admin.Get("/user/:id", controllers.ShowUser)
	// tag
	//admin.Get("/account/new", controllers.NewAccount)
	//admin.Post("/account", controllers.CreateAccount)

	log.Fatal(app.Listen(":8080"))
}
