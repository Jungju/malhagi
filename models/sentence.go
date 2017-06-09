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
	Id          int64        `json:"id"`
	CreatedAt   time.Time    `json:"created_at" orm:"type(datetime);auto_now"`
	Text        string       `json:"text",orm:"size(128)"`
	Korean      string       `json:"korean",orm:"size(128)"`
	VerbsType   verbs.Type   `json:"verbsType"`
	PersonsType persons.Type `json:"personsType"`
	FormatsType formats.Type `json:"formatsType"`
	TensesType  tenses.Type  `json:"tensesType"`
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
		if !game.CheckFormatsType(sentence.FormatsType) {
			continue
		}
		if !game.CheckVerbType(sentence.VerbsType) {
			continue
		}
		if !game.CheckPersonsType(sentence.PersonsType) {
			continue
		}
		if !game.CheckTensesType(sentence.TensesType) {
			continue
		}
		filterdSentences = append(filterdSentences, &sentence)
	}

	if len(filterdSentences) == 0 {
		return nil
	}

	return filterdSentences[rand.Intn(len(filterdSentences))]
}
