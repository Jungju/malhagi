package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
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

// TestGet is a sample to run an endpoint test
func TestSentencePost(t *testing.T) {
	Convey("단어 만들기", t, func() {
		Convey("단어 만들기 성공", func() {
			reqTest("POST", "/sentence", models.Sentence{
				Text:        "I am sleep.",
				Korean:      "나는 잔다",
				TensesType:  tenses.Present,
				VerbsType:   verbs.BeVerb,
				PersonsType: persons.I,
				FormatsType: formats.Plain,
			}, 201, "")
		})
		Convey("단어 확인", func() {
			bodyBytes := reqTest("GET", "/sentence", nil, 200, "")
			sentences := []models.Sentence{}
			err := json.Unmarshal(bodyBytes, &sentences)
			So(err, ShouldBeNil)
			So(len(sentences), ShouldEqual, 1)
		})
	})
}

func reqTest(method string, url string, body interface{}, expectedCode int, token string) []byte {
	bodyBytes, _ := json.Marshal(body)
	r, _ := http.NewRequest(method, fmt.Sprintf("/%s%s", APIVersion, url), bytes.NewBuffer(bodyBytes))
	if token != "" {
		r.Header.Add("X-Auth-Token", token)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("BODY : %s", w.Body.String())
	So(w.Code, ShouldEqual, expectedCode)

	return w.Body.Bytes()
}

//TODO: 단어 리스트

//TODO: 단어 수정

//TODO: 단어 삭제
