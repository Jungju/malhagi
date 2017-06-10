package models

import (
	"time"

	"strings"

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
	VerbsTypes   string    `json:"verbsTypes"`
	PersonsTypes string    `json:"personsTypes"`
	FormatsTypes string    `json:"formatsTypes"`
	TensesTypes  string    `json:"tensesTypes"`
}

//GetAllVerbTypeIds ...
func (m Game) GetAllVerbTypeIds() []int {
	if m.VerbsTypes == "" {
		return verbs.Ids()
	}

	ids := []int{}
	types := strings.Split(m.VerbsTypes, ",")
	for _, stringType := range types {
		if verbs.Be.String() == stringType {
			ids = append(ids, int(verbs.Be))
		} else if verbs.General.String() == stringType {
			ids = append(ids, int(verbs.Be))
		}
	}
	return ids
}

//GetAllFormatTypeIds ...
func (m Game) GetAllFormatTypeIds() []int {
	if m.FormatsTypes == "" {
		return formats.Ids()
	}

	ids := []int{}
	types := strings.Split(m.FormatsTypes, ",")
	for _, stringType := range types {
		if formats.Plain.String() == stringType {
			ids = append(ids, int(formats.Plain))
		} else if formats.Future.String() == stringType {
			ids = append(ids, int(formats.Future))
		} else if formats.Question.String() == stringType {
			ids = append(ids, int(formats.Question))
		} else if formats.Negative.String() == stringType {
			ids = append(ids, int(formats.Negative))
		}
	}
	return ids
}

//GetAllTensesTypeIds ...
func (m Game) GetAllTensesTypeIds() []int {
	if m.TensesTypes == "" {
		return tenses.Ids()
	}

	ids := []int{}
	types := strings.Split(m.TensesTypes, ",")
	for _, stringType := range types {
		if tenses.Past.String() == stringType {
			ids = append(ids, int(tenses.Past))
		} else if tenses.Present.String() == stringType {
			ids = append(ids, int(tenses.Present))
		} else if tenses.Future.String() == stringType {
			ids = append(ids, int(tenses.Future))
		}
	}
	return ids
}

//GetAllPersonsTypeIds ...
func (m Game) GetAllPersonsTypeIds() []int {
	if m.PersonsTypes == "" {
		return persons.Ids()
	}

	ids := []int{}
	types := strings.Split(m.PersonsTypes, ",")
	for _, stringType := range types {
		if persons.I.String() == stringType {
			ids = append(ids, int(persons.I))
		} else if persons.We.String() == stringType {
			ids = append(ids, int(persons.We))
		} else if persons.You.String() == stringType {
			ids = append(ids, int(persons.You))
		} else if persons.They.String() == stringType {
			ids = append(ids, int(persons.They))
		} else if persons.He.String() == stringType {
			ids = append(ids, int(persons.He))
		} else if persons.She.String() == stringType {
			ids = append(ids, int(persons.She))
		} else if persons.It.String() == stringType {
			ids = append(ids, int(persons.It))
		} else if persons.Special.String() == stringType {
			ids = append(ids, int(persons.Special))
		}
	}
	return ids
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
