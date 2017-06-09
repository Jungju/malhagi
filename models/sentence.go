package models

import (
	"math/rand"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/jungju/malhagi/types/formats"
	"github.com/jungju/malhagi/types/persons"
	"github.com/jungju/malhagi/types/tenses"
	"github.com/jungju/malhagi/types/verbs"
)

var GlobalSentence []Sentence

//Sentence ...
//bee generate model sentence -fields="id:int,created_at:datetime,text:string,korean:string,verbs_type:int,persons_type:int,formats_type:int,tenses_type:int"
type Sentence struct {
	Id          int64
	CreatedAt   time.Time `orm:"type(datetime);auto_now"`
	Text        string    `orm:"size(128)"`
	Korean      string    `orm:"size(128)"`
	VerbsType   verbs.Type
	PersonsType persons.Type
	FormatsType formats.Type
	TensesType  tenses.Type
}

//ValidCreate ...
func (m Sentence) ValidCreate() bool {
	if m.Text == "" || m.Korean == "" {
		return false
	}
	return true
}

//ValidUpdate ...
func (m Sentence) ValidUpdate() bool {
	if m.Id <= 0 {
		return false
	}
	if m.Text == "" || m.Korean == "" {
		return false
	}
	return true
}

// AddSentence ...
func AddSentence(m *Sentence) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(m)
}

// GetSentenceById ...
func GetSentenceById(id int64) (*Sentence, error) {
	o := orm.NewOrm()
	v := &Sentence{Id: id}
	err := o.Read(v)
	return v, err
}

// GetSentenceByText ...
func GetSentenceByText(text string) (*Sentence, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sentence))
	qs = qs.Filter("text", text)

	sentence := &Sentence{}
	err := qs.One(sentence)
	return sentence, err
}

// GetAllSentence ...
func GetAllSentence() ([]Sentence, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sentence))
	qs = qs.Limit(500)
	var l []Sentence
	_, err := qs.All(&l)
	return l, err
}

// UpdateSentence ...
func UpdateSentence(m *Sentence) error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}

// DeleteSentence ...
func DeleteSentence(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Sentence{Id: id})
	return err
}

// GetRandomSentence ...
func GetRandomSentence(game *Game) *Sentence {
	if GlobalSentence == nil && len(GlobalSentence) == 0 {
		GlobalSentence, _ = GetAllSentence()
	}

	allSentence := GlobalSentence

	filterdSentences := []*Sentence{}
	for _, sentence := range allSentence {
		if game.FormatsType != formats.None && sentence.FormatsType != game.FormatsType {
			continue
		}
		if game.VerbsType != verbs.None && sentence.VerbsType != game.VerbsType {
			continue
		}
		if game.PersonsType != persons.None && sentence.PersonsType != game.PersonsType {
			continue
		}
		if game.TensesType != tenses.None && sentence.TensesType != game.TensesType {
			continue
		}
		filterdSentences = append(filterdSentences, &sentence)
	}

	if len(filterdSentences) == 0 {
		return nil
	}

	return filterdSentences[rand.Intn(len(filterdSentences))]
}
