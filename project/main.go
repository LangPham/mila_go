package main

import (
	"github.com/LangPham/mila/apps_web"
	"github.com/LangPham/mila/config"
	debug "github.com/LangPham/mila_debug"
)

func main() {
	debug.Dump(config.Config.Get("view.dir"), "MAIN")

	apps_web.AppsWeb()
}
