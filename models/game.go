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
	Id           int64     `json:"id"`
	CreatedAt    time.Time `json:"created_at" orm:"type(datetime);auto_now"`
	Ended        bool      `json:"ended"`
	Point        int       `json:"point"`
	VerbsTypes   int       `json:"verbsTypes"`
	PersonsTypes int       `json:"personsTypes"`
	FormatsTypes int       `json:"formatsTypes"`
	TensesTypes  int       `json:"tensesTypes"`
}

//CheckVerbType ...
func (m Game) CheckVerbType(t verbs.Type) bool {
	return m.VerbsTypes == 0 || m.VerbsTypes&int(t) == int(t)
}

//CheckPersonsType ...
func (m Game) CheckPersonsType(t persons.Type) bool {
	return m.PersonsTypes == 0 || m.PersonsTypes&int(t) == int(t)
}

//CheckFormatsType ...
func (m Game) CheckFormatsType(t formats.Type) bool {
	return m.FormatsTypes == 0 || m.FormatsTypes&int(t) == int(t)
}

//CheckTensesType ...
func (m Game) CheckTensesType(t tenses.Type) bool {
	return m.TensesTypes == 0 || m.TensesTypes&int(t) == int(t)
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
