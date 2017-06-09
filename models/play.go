package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Play
//bee generate model Play -fields="id:int,created_at:datetime,ended:bool,point:int"
type Play struct {
	Id         int64
	CreatedAt  time.Time `orm:"type(datetime);auto_now"`
	SentenceID int64
	GameId     int64
	Input      string
	IsSuccess  bool
}

//ValidCreate ...
func (m Play) ValidCreate() bool {
	return true
}

//ValidUpdate ...
func (m Play) ValidUpdate() bool {
	return true
}

// AddPlay ...
func AddPlay(m *Play) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetPlayById ...
func GetPlayById(id int64) (*Play, error) {
	o := orm.NewOrm()
	v := &Play{Id: id}
	err := o.Read(v)
	return v, err
}

// GetAllPlay ...
func GetAllPlay() ([]Play, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Play))

	var l []Play
	_, err := qs.All(&l)
	return l, err
}

// UpdatePlay
func UpdatePlay(m *Play) error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}

func GetPlaySuccessCountByGameID(GameId int64) (int, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Play))
	qs = qs.Filter("game_id", GameId)
	qs = qs.Filter("is_success", true)

	cnt, err := qs.Count()
	return int(cnt), err
}
