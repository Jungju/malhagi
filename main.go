package main

import (
	"fmt"
	"os"
	"time"

	"github.com/boltdb/bolt"
	"github.com/jungju/malhagi/models"
	_ "github.com/jungju/malhagi/routers"

	"github.com/astaxie/beego"
)

var (
	dbName string
	port   string
	host   string
)

func init() {
	dbName = os.Getenv("DB_NAME")
	port = os.Getenv("WEB_PORT")
	host = os.Getenv("HOST")
	if port == "" {
		port = "8373"
	}
	if host == "" {
		host = "localhost"
	}
	if dbName == "" {
		dbName = "db/malhagi.db"
	}
}

func main() {
	var err error
	if models.DB, err = bolt.Open(dbName, 0600, &bolt.Options{Timeout: 2 * time.Second}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
