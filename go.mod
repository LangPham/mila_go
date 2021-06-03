module github.com/LangPham/mila_go

go 1.16

replace github.com/LangPham/mila_cast => /home/langpham/git/mila_cast

replace github.com/LangPham/mila_guard => /home/langpham/git/mila_guard

require (
	github.com/LangPham/mila_cast v0.0.0-00010101000000-000000000000
	github.com/LangPham/mila_guard v0.0.0-00010101000000-000000000000
	github.com/aymerick/raymond v2.0.2+incompatible
	github.com/casbin/casbin/v2 v2.31.2
	github.com/davecgh/go-spew v1.1.1
	github.com/gofiber/fiber/v2 v2.10.0
	github.com/lib/pq v1.8.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/valyala/fasthttp v1.24.0 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.9
)
