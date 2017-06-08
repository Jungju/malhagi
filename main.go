package main

import (
	"github.com/jungju/malhagi/models"
	_ "github.com/jungju/malhagi/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	if err := models.InitDB(); err != nil {
		panic(err)
	}
	beego.Run()
}
