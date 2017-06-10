package models

import (
	"math/rand"
	"time"

	"fmt"

	"strings"

	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/jungju/malhagi/types/formats"
	"github.com/jungju/malhagi/types/persons"
	"github.com/jungju/malhagi/types/tenses"
	"github.com/jungju/malhagi/types/verbs"
)

//Sentence ...
//bee generate model sentence -fields="id:int,created_at:datetime,text:string,korean:string,verbs_type:int,persons_type:int,formats_type:int,tenses_type:int"
type Sentence struct {
	Id          int64        `json:"id"`
	CreatedAt   time.Time    `json:"created_at" orm:"type(datetime);auto_now"`
	Text        string       `json:"text,omitempty",orm:"size(128)"`
	Korean      string       `json:"korean,omitempty",orm:"size(128)"`
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
	if m.FormatsType == formats.None ||
		m.TensesType == tenses.None ||
		m.PersonsType == persons.None ||
		m.VerbsType == verbs.None {
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
func GetRandomSentence(game *Game) (*Sentence, error) {
	o := orm.NewOrm()
	var sentences []Sentence
	sql := "SELECT sentence.* FROM sentence WHERE 1=1 "
	sql += fmt.Sprintf("AND verbs_type IN (%s) ", intArrayJoin(game.GetAllVerbTypeIds()))
	sql += fmt.Sprintf("AND formats_type IN (%s) ", intArrayJoin(game.GetAllFormatTypeIds()))
	sql += fmt.Sprintf("AND tenses_type IN (%s) ", intArrayJoin(game.GetAllTensesTypeIds()))
	sql += fmt.Sprintf("AND persons_type IN (%s) ", intArrayJoin(game.GetAllPersonsTypeIds()))
	sql += fmt.Sprintf("AND sentence.id NOT IN (SELECT sentence_id FROM play WHERE game_id = %d) ", game.Id)
	_, err := o.Raw(sql).QueryRows(&sentences)
	if err != nil {
		return nil, err
	}

	if len(sentences) == 0 {
		return nil, errors.New("준비 된 문제 없음")
	}
	sentence := sentences[rand.Intn(len(sentences))]

	return &sentence, err
}

func intArrayJoin(integers []int) string {
	strArrat := []string{}
	for _, i := range integers {
		strArrat = append(strArrat, fmt.Sprintf("%d", i))
	}
	return strings.Join(strArrat, ",")
}
