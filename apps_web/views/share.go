package views

import (
	"github.com/aymerick/raymond"
	"github.com/gofiber/fiber/v2"
	"path/filepath"
)

type M map[string]interface{}

func InnerContent(file string, m M) (result string) {

	result = ""
	fullFile := file + ".hgo"
	fileTpl := filepath.Join("apps_web/templates", fullFile)
	tpl, e := raymond.ParseFile(fileTpl)

	if e != nil {
		return
	}
	result = tpl.MustExec(m)
	return
}

func RenderHTML(c *fiber.Ctx, file string, m M, option ...string) error {
	masterFile := "apps_web/templates/layouts/master.hgo"
	masterTpl, _ := raymond.ParseFile(masterFile)

	masterTpl.RegisterHelper("inner_content", func() raymond.SafeString {
		return raymond.SafeString(InnerContent(file, m))
	})
	//debug.Dump(masterTpl, "MASTER")
	resultMaster := masterTpl.MustExec(m)
	c.Set("Content-Type", "text/html; charset=utf8")
	return c.SendString(resultMaster)
	//c.SetContentType("text/html; charset=utf8")
	//c.SetBodyString(resultMaster)
}

//func LiveInner(file string, m M) (result string) {
//
//	result = ""
//	fullFile := file + ".hgo"
//	fileTpl := filepath.Join("apps_web/templates", fullFile)
//	var AppFs = afero.NewOsFs()
//	tplByte, _ := afero.ReadFile(AppFs, fileTpl)
//	tpl, _ := raymond.Parse(string(tplByte))
//	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(tplByte)))
//	sec := doc.Find("[eta-update]")
//	id, _ := sec.Attr("id")
//	strHTML, _ :=  sec.Html()
//	before := `<script id="template_` + id  + `" type="text/x-handlebars-template">`
//	end := `</script>`
//	strHTML = before+strHTML+end
//	result = strHTML+ tpl.MustExec(m)//string(tplByte)
//	return
//}

//func LiveRender(ctx *fasthttp.RequestCtx, file string, m M, option ...string) {
//	masterFile := "apps_web/templates/layouts/live.hgo"
//	masterTpl, eMaster := raymond.ParseFile(masterFile)
//	if eMaster != nil {
//		fmt.Println("ERROR: " + eMaster.Error())
//	}
//	masterTpl.RegisterHelper("inner_content", func() raymond.SafeString {
//		return raymond.SafeString(LiveInner(file, m))
//	})
//	resultMaster := masterTpl.MustExec(m)
//	ctx.SetContentType("text/html; charset=utf8")
//	ctx.SetBodyString(resultMaster)
//}
