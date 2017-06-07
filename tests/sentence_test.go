package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jungju/malhagi/models"
	_ "github.com/jungju/malhagi/routers"

	"encoding/json"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

//TODO: 단어 만들기
// TestGet is a sample to run an endpoint test
func TestSentencePost(t *testing.T) {
	bodyBytes, _ := json.Marshal(&models.Sentence{
		Verb:   "sleep",
		Korean: "자다",
		//과거 필요 : 잤었다.
	})
	r, _ := http.NewRequest("POST", "/v1/sentence", bytes.NewBuffer(bodyBytes))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("단어 만들기", t, func() {
		Convey("단어 만들기 성공", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("단어 확인", func() {
			So(w.Code, ShouldNotImplement)
		})
	})
}

//TODO: 단어 리스트

//TODO: 단어 수정

//TODO: 단어 삭제
