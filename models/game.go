package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/jungju/malhagi/types/formats"
	"github.com/jungju/malhagi/types/persons"
	"github.com/jungju/malhagi/types/tenses"
	"github.com/jungju/malhagi/types/verbs"
)

//Game
//bee generate model game -fields="id:int,created_at:datetime,ended:bool,point:int"
type Game struct {
	Id          int64
	CreatedAt   time.Time `orm:"type(datetime);auto_now"`
	Ended       bool
	Point       int
	VerbsType   verbs.Type
	PersonsType persons.Type
	FormatsType formats.Type
	TensesType  tenses.Type
}

//ValidCreate ...
func (m Game) ValidCreate() bool {
	return true
}

//ValidUpdate ...
func (m Game) ValidUpdate() bool {
	return true
}

// AddGame ...
func AddGame(m *Game) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetGameById ...
func GetGameById(id int64) (*Game, error) {
	o := orm.NewOrm()
	v := &Game{Id: id}
	err := o.Read(v)
	return v, err
}

// GetAllGame ...
func GetAllGame() ([]Game, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Game))
	qs = qs.OrderBy("-point")
	qs = qs.Limit(100)
	var l []Game
	_, err := qs.All(&l)
	return l, err
}

// UpdateGame
func UpdateGame(m *Game) error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}
