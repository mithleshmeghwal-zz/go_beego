package main

import (
	_ "ormTest/models"
	_ "ormTest/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1234@tcp(127.0.0.1:3306)/test")
	// 	orm.RunSyncdb(
	// 		"default",
	// 		true,
	// 		true,
	// 	)
}

func main() {
	orm.Debug = true
	beego.Run()
}
