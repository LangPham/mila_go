package repo

import (
	"log"
	"os"

	"github.com/LangPham/mila_go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"database/sql"
    _ "github.com/lib/pq"
    "github.com/golang-migrate/migrate/v4"
    pg "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

var Repo *gorm.DB

func init() {
	theta_migrate()
	//a := config.GetJoin("repo")
	//aon.Dump(a)
	//dsn := "host=127.0.0.1 user=postgres password=postgres dbname=mila port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	dsn := config.GetJoin("repo")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("ERROR: fail create postgres session, %s", err.Error())
		os.Exit(1)
	}
	Repo = db
	//a, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
	//aon.Dump(a, "CASBIN")// util.Dump(dsn)
//r = sub, obj, act
//
//[policy_definition]
//p = sub, obj, act
//
//[role_definition]
//g = _, _
//
//[policy_effect]
//e = some(where (p.eft == allow))
//
//[matchers]
//m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
//
//`)
//	aon.Dump(m, "CASBIN")
	//e, _ := casbin.NewEnforcer(m, a)

	//e.LoadPolicy()
	//CasEnf = e
	//aon.Dump(e, "CASBIN")
}

func theta_migrate() { 
    dri, conf := config.GetRepo("repo")
    // util.Dump(dsn)

	db, _ := sql.Open(dri, conf)
    driver, _ := pg.WithInstance(db, &pg.Config{})
    m, _ := migrate.NewWithDatabaseInstance(
        "file://./apps/db/migrations",
        "postgres", driver)
    m.Up()
}
