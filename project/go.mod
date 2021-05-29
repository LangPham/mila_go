module github.com/LangPham/mila

go 1.16

replace github.com/LangPham/mila_view => /Volumes/Master/Git/mila_go/mila_view

replace github.com/LangPham/mila_debug => /Volumes/Master/Git/mila_go/mila_debug

//replace github.com/LangPham/mila_controller => /Volumes/Master/Git/mila_go/mila_controller

require (
	github.com/LangPham/mila_debug v0.0.0-00010101000000-000000000000
	github.com/LangPham/mila_view v0.0.0-00010101000000-000000000000
	github.com/aymerick/raymond v2.0.2+incompatible
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/gofiber/fiber/v2 v2.9.0
	github.com/spf13/viper v1.7.1
	github.com/valyala/fasthttp v1.24.0 // indirect
	gorm.io/driver/postgres v1.1.0 // indirect
	gorm.io/gorm v1.21.9 // indirect
)
