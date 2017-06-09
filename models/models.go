package models

import (
	"fmt"

	"github.com/jungju/malhagi/envs"

	"github.com/astaxie/beego/orm"
)

//InitDB ...
func InitDB() error {
	orm.Debug = false

	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", envs.DBUser, envs.DBPassword, envs.DBHost, envs.DBDatabase)); err != nil {
		return err
	}
	orm.RegisterModel(
		new(Sentence),
		new(Game),
		new(Play),
	)

	return orm.RunSyncdb("default", true, true)
}
