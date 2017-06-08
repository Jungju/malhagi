package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/jungju/malhagi/types/formats"
	"github.com/jungju/malhagi/types/persons"
	"github.com/jungju/malhagi/types/tenses"
	"github.com/jungju/malhagi/types/verbs"
)

//Sentence ...
//bee generate model sentence -fields="id:int,created_at:datetime,text:string,korean:string,verbs_type:int,persons_type:int,formats_type:int,tenses_type:int"
type Sentence struct {
	Id          int64
	CreatedAt   time.Time `orm:"type(datetime);auto_now)"`
	Text        string    `orm:"size(128)"`
	Korean      string    `orm:"size(128)"`
	VerbsType   verbs.Type
	PersonsType persons.Type
	FormatsType formats.Type
	TensesType  tenses.Type
}

// AddSentence insert a new Sentence into database and returns
// last inserted Id on success.
func AddSentence(m *Sentence) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSentenceById retrieves Sentence by Id. Returns error if
// Id doesn't exist
func GetSentenceById(id int64) (v *Sentence, err error) {
	o := orm.NewOrm()
	v = &Sentence{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSentence retrieves all Sentence matches certain condition. Returns empty list if
// no records exist
func GetAllSentence() ([]Sentence, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sentence))

	var l []Sentence
	_, err := qs.All(&l)
	return l, err
}

// UpdateSentence updates Sentence by Id and returns error if
// the record to be updated doesn't exist
func UpdateSentenceById(m *Sentence) (err error) {
	o := orm.NewOrm()
	v := Sentence{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSentence deletes Sentence by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSentence(id int64) (err error) {
	o := orm.NewOrm()
	v := Sentence{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Sentence{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
