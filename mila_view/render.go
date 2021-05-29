package mila_view

import (
	dbug "github.com/LangPham/mila_debug"
	"github.com/aymerick/raymond"
	"github.com/gofiber/fiber/v2"
	"path"
)

func New(dirTem string, layout string, ext string) *View {
	fileFull := layout + "." + ext
	fileLayout := path.Join(dirTem, fileFull)
	masterTpl, _ := raymond.ParseFile(fileLayout)
	view := &View{
		DirTemplate: dirTem,
		Layout:      layout,
		Extension:   ext,
		TemplateMaster: masterTpl,
	}

	return view
}

func (view View) InnerContent(mtpl map[string]interface{}, m map[string]interface{}) (result string) {
	fileTpl := view.getFile(mtpl)
	tpl, e := raymond.ParseFile(fileTpl)

	if e != nil {
		return
	}
	result = tpl.MustExec(m)
	return
}

func (view *View) HTML(c *fiber.Ctx, tpl map[string]interface{}, m map[string]interface{}) error {
	//
	//if tpl["layout"] != nil {
	//	masterFile := view.getLayout(tpl)
	//	//masterTpl, eMaster := raymond.ParseFile(masterFile)
	//	masterTpl, _ := raymond.ParseFile(masterFile)
	//	view.TemplateMaster = masterTpl
	//}

	//templateInner := view.TemplateMaster.Clone()
	//dbug.Dump(view.TemplateMaster, "T MASTER")
	view.TemplateMaster.RegisterHelper("inner_content", func() raymond.SafeString {
		return raymond.SafeString(view.InnerContent(tpl, m))
	})
	dbug.Dump(view.TemplateMaster, "T INNER")
	resultMaster := view.TemplateMaster.MustExec(m)
	//resultMaster := "a"

	c.Set("Content-Type", "text/html; charset=utf8")
	return c.SendString(resultMaster)
}


// Return file layout with folder
func (view View) getLayout(option map[string]interface{}) string {
	layout := view.Layout
	if option["layout"] != nil {
		layout = option["layout"].(string)
	}
	layoutFull := view.fullFile(layout)
	fileLayout := path.Join(view.DirTemplate, layoutFull)
	return fileLayout
}

// Return file tpl with folder
func (view View) getFile(option map[string]interface{}) string {
	fileName := option["tpl"].(string)
	return path.Join(view.DirTemplate, view.fullFile(fileName))
}

// Add Extension for name to full name file
func (view View) fullFile(name string) string {
	return name + "." + view.Extension
}

