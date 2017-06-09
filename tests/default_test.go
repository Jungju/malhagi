package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"

	"github.com/jungju/malhagi/models"

	_ "github.com/jungju/malhagi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	APIVersion = "v1"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	models.InitDB()
	orm.Debug = true
}

func reqTest(method string, url string, body interface{}, expectedCode int, token string, testing bool) []byte {
	var r *http.Request
	if body == nil {
		r, _ = http.NewRequest(method, fmt.Sprintf("/%s%s", APIVersion, url), nil)
	} else {
		bodyBytes, _ := json.Marshal(body)
		r, _ = http.NewRequest(method, fmt.Sprintf("/%s%s", APIVersion, url), bytes.NewBuffer(bodyBytes))
	}

	if token != "" {
		r.Header.Add("X-Auth-Token", token)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	beego.Trace("BODY : %s", w.Body.String())
	if testing {
		So(w.Code, ShouldEqual, expectedCode)
	}

	return w.Body.Bytes()
}
