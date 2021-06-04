package apps_web

import (
	"github.com/LangPham/mila_go/util"
	guard "github.com/LangPham/mila_guard"
	"log"

	"github.com/LangPham/mila_go/apps_web/controllers"
	"github.com/LangPham/mila_go/config"
	"github.com/gofiber/fiber/v2"
)

func AppsWeb() {
	app := fiber.New()
	app.Static("/", "./public/static")
	//app.Get("/", controllers.PageIndex)
	app.Use(func(c *fiber.Ctx) error {
		//aon.Dump(c.Method())
		//aon.Dump(c.Path(), "PATH")
		//r := c.Route()
		//aon.Dump(r, "ROUTE")

		switch c.Method() {
		case "POST":
			if c.FormValue("_METHOD") != "" {
				c.Method(c.FormValue("_METHOD"))
			}
		default:

		}
		// Set some security headers:
		//c.Set("X-XSS-Protection", "1; mode=block")
		//c.Set("X-Content-Type-Options", "nosniff")
		//c.Set("X-Download-Options", "noopen")
		//c.Set("Strict-Transport-Security", "max-age=5184000")
		//c.Set("X-Frame-Options", "SAMEORIGIN")
		//c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		util.Dump(c.Path(), "PATH")
		cookieName := config.Config.Get("cookie.name").(string)
		cookieValue := c.Cookies(cookieName)
		var gua guard.Guard
		if cookieValue != "" {
			// Have cookie
			strGuard := util.MilaDecrypt(cookieValue)
			gua = guard.ToGuard(strGuard)
		} else {
			cookie := new(fiber.Cookie)
			cookie.HTTPOnly = true
			cookie.Name = cookieName
			gua = guard.Guard{Role: "GUEST", UserID: ""}
			strGuard := gua.ToString()
			cookie.Value = util.MilaEncrypt(strGuard)
			c.Cookie(cookie)
		}
		//util.Dump(gua, "GUA EXIST")
		enforce, err := Cas.Enforce(gua.Role, c.Path(), c.Method())
		//util.Dump(enforce, "enforce")
		util.Dump(err, "err")
		//util.Dump(cookieValue, "cookieName")
		//Create cookie session
		//cookie := new(fiber.Cookie)
		//cookie.HTTPOnly = true
		//cookie.Name = "name"
		//cookie.Value = "doe"
		// Set cookie
		//c.Cookie(cookie)

		//aon.Dump(c.Cookies("name")  ,"NAME"  )     // "john"
		//aon.Dump(c.Cookies("empty", "doe"), "2 key") // "doe"
		if enforce {
			return c.Next()
		} else {
			return c.SendStatus(403)
		}

		//return c.SendString("USE")
	})
	app.Get("/", controllers.Test, controllers.Test1)
	app.Get("/login", controllers.Login)

	admin := app.Group("/admin")
	// account
	account := admin.Group("/account")
	account.Get("/new", controllers.NewAccount)

	account.Get("/:id", controllers.ShowAccount)
	account.Get("/:id/edit", controllers.EditAccount)
	account.Post("/", controllers.CreateAccount)
	account.Put("/:id", controllers.UpdateAccount)

	tag:= admin.Group("tag")
	tag.Get("/new", controllers.NewTag)
	tag.Post("/", controllers.CreateTag)
	//admin.Get("/user/:id", controllers.ShowUser)
	// tag
	//admin.Get("/account/new", controllers.NewAccount)
	//admin.Post("/account", controllers.CreateAccount)
	//aon.Dump(app.Stack(), "Stack")
	//data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	//fmt.Println(string(data))
	log.Fatal(app.Listen(":8080"))
}
