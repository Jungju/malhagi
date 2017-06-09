package test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jungju/malhagi/models"
	_ "github.com/jungju/malhagi/routers"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/jungju/malhagi/types/formats"
	"github.com/jungju/malhagi/types/persons"
	"github.com/jungju/malhagi/types/tenses"
	"github.com/jungju/malhagi/types/verbs"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestSentencePost(t *testing.T) {
	Convey("단어 만들기", t, func() {
		Convey("단어 만들기 실패", func() {
			reqTest("POST", "/sentence", models.Sentence{
				Korean:      "나는 잔다",
				TensesType:  tenses.Present,
				VerbsType:   verbs.BeVerb,
				PersonsType: persons.I,
				FormatsType: formats.Plain,
			}, 400, "", true)
		})
		Convey("단어 만들기 성공", func() {
			bodyBytes := reqTest("GET", "/sentence", nil, 200, "", true)
			sentences := []models.Sentence{}
			err := json.Unmarshal(bodyBytes, &sentences)
			cnt := len(sentences)

			reqTest("POST", "/sentence", models.Sentence{
				Text:        "I am sleep",
				Korean:      "나는 잔다",
				TensesType:  tenses.Present,
				VerbsType:   verbs.BeVerb,
				PersonsType: persons.I,
				FormatsType: formats.Plain,
			}, 201, "", true)

			nextBodyBytes := reqTest("GET", "/sentence", nil, 200, "", true)
			nextSentences := []models.Sentence{}
			err = json.Unmarshal(nextBodyBytes, &nextSentences)
			So(err, ShouldBeNil)
			So(len(nextSentences), ShouldEqual, cnt+1)
		})
		Convey("중복된 단어는 안됨", func() {
			reqTest("POST", "/sentence", models.Sentence{
				Text:        "I am sleep",
				Korean:      "나는 잔다2",
				TensesType:  tenses.Present,
				VerbsType:   verbs.BeVerb,
				PersonsType: persons.I,
				FormatsType: formats.Plain,
			}, 409, "", true)
		})
	})
}

func TestSentencePut(t *testing.T) {
	Convey("단어 수정하기", t, func() {
		Convey("단어 수정하기 실패", func() {
			reqTest("PUT", "/sentence/1", models.Sentence{
				Text:        "",
				Korean:      "나는 잔다2",
				TensesType:  tenses.Present,
				VerbsType:   verbs.BeVerb,
				PersonsType: persons.I,
				FormatsType: formats.Plain,
			}, 400, "", true)
			reqTest("PUT", "/sentence/10", models.Sentence{
				Id:          10,
				Text:        "I am sleep.2",
				Korean:      "나는 잔다2",
				TensesType:  tenses.Present,
				VerbsType:   verbs.BeVerb,
				PersonsType: persons.I,
				FormatsType: formats.Plain,
			}, 404, "", true)
		})
		Convey("단어 수정하기 성공", func() {
			reqTest("PUT", "/sentence/1", models.Sentence{
				Id:          1,
				Text:        "I am sleep.2",
				Korean:      "나는 잔다2",
				TensesType:  tenses.Present,
				VerbsType:   verbs.BeVerb,
				PersonsType: persons.I,
				FormatsType: formats.Plain,
			}, 204, "", true)

			bodyBytes := reqTest("GET", "/sentence", nil, 200, "", true)
			sentences := []models.Sentence{}
			err := json.Unmarshal(bodyBytes, &sentences)
			So(err, ShouldBeNil)
			So(sentences[0].Text, ShouldEqual, "I am sleep.2")
		})
	})
}

func TestSentenceDelete(t *testing.T) {
	Convey("단어 삭제하기", t, func() {
		Convey("단어 삭제하기 실패", func() {
			reqTest("DELETE", "/sentence/10", nil, 404, "", true)
		})
		Convey("단어 삭제하기 성공", func() {
			bodyBytes := reqTest("GET", "/sentence", nil, 200, "", true)
			sentences := []models.Sentence{}
			err := json.Unmarshal(bodyBytes, &sentences)
			cnt := len(sentences)

			reqTest("DELETE", "/sentence/1", nil, 204, "", true)

			nextBodyBytes := reqTest("GET", "/sentence", nil, 200, "", true)
			nextSentences := []models.Sentence{}
			err = json.Unmarshal(nextBodyBytes, &nextSentences)
			So(err, ShouldBeNil)
			So(len(nextSentences), ShouldEqual, cnt-1)
		})
	})
}
