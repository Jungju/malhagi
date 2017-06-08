package test

import (
	"path/filepath"
	"runtime"

	"github.com/jungju/malhagi/models"

	_ "github.com/jungju/malhagi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	APIVersion = "v1"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	models.InitDB()
	orm.Debug = true
}
