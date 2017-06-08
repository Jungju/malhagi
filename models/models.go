package models

import (
	"fmt"

	"bitbucket.org/jungju/ahg/envs"

	"github.com/astaxie/beego/orm"
)

//InitDB ...
func InitDB() error {
	orm.Debug = true

	orm.RegisterDriver("mysql", orm.DRMySQL)
	if err := orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", envs.DBUser, envs.DBPassword, envs.DBHost, envs.DBDatabase)); err != nil {
		return err
	}
	orm.RegisterModel(
		new(Sentence),
		new(Game),
	)

	return orm.RunSyncdb("default", true, true)
}
