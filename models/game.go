package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

//Game
//bee generate model game -fields="id:int,created_at:datetime,ended:bool,point:int"
type Game struct {
	Id        int64
	CreatedAt time.Time `orm:"type(datetime);auto_now)"`
	Ended     bool
	Point     int
}

// AddGame insert a new Game into database and returns
// last inserted Id on success.
func AddGame(m *Game) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGameById retrieves Game by Id. Returns error if
// Id doesn't exist
func GetGameById(id int64) (v *Game, err error) {
	o := orm.NewOrm()
	v = &Game{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGame retrieves all Game matches certain condition. Returns empty list if
// no records exist
func GetAllGame() ([]Game, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Game))

	var l []Game
	_, err := qs.All(&l)
	return nil, err
}

// UpdateGame updates Game by Id and returns error if
// the record to be updated doesn't exist
func UpdateGameById(m *Game) (err error) {
	o := orm.NewOrm()
	v := Game{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
