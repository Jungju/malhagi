package test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/jungju/malhagi/models"
	_ "github.com/jungju/malhagi/routers"
	"github.com/jungju/malhagi/types/formats"
	"github.com/jungju/malhagi/types/persons"
	"github.com/jungju/malhagi/types/tenses"
	"github.com/jungju/malhagi/types/verbs"

	"encoding/json"

	"fmt"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestGamePost(t *testing.T) {
	Convey("게임 만들기", t, func() {
		Convey("게임 만들기 성공", func() {
			reqTest("POST", "/game", nil, 201, "", true)
		})
		Convey("게임 확인", func() {
			bodyBytes := reqTest("GET", "/game/1", nil, 200, "", true)
			game := models.Game{}
			err := json.Unmarshal(bodyBytes, &game)
			So(err, ShouldBeNil)
			So(game.Ended, ShouldEqual, false)
		})
	})
}

func TestGamePut(t *testing.T) {
	Convey("게임 종료", t, func() {
		Convey("게임 종료 성공", func() {
			bodyBytes := reqTest("PUT", "/game/1/end", nil, 200, "", true)
			game := models.Game{}
			err := json.Unmarshal(bodyBytes, &game)
			So(err, ShouldBeNil)
			So(game.Ended, ShouldEqual, true)
		})
		Convey("게임 확인", func() {
			bodyBytes := reqTest("GET", "/game/1", nil, 200, "", true)
			game := models.Game{}
			err := json.Unmarshal(bodyBytes, &game)
			So(err, ShouldBeNil)
			So(game.Ended, ShouldEqual, true)
		})
	})
}

func TestGamePlay(t *testing.T) {
	game := models.Game{
	// TensesTypes:  "past",
	// FormatsTypes: "plain",
	// VerbsTypes:   "be",
	// PersonsTypes: "i",
	}
	gameBodyBytes := reqTest("POST", "/game", &game, 201, "", false)
	json.Unmarshal(gameBodyBytes, &game)
	gameID := game.Id

	sentenceBytes := reqTest("POST", "/sentence", models.Sentence{
		Text:        "I study",
		Korean:      "나는 공부한다",
		TensesType:  tenses.Present,
		VerbsType:   verbs.General,
		PersonsType: persons.I,
		FormatsType: formats.Plain,
	}, 201, "", false)
	sentence := models.Sentence{}
	json.Unmarshal(sentenceBytes, &sentence)
	sentenceID := sentence.Id

	Convey("문제 풀기", t, func() {
		//TODO: 문제 받고 풀기로 변경
		Convey("문제 받기 성공", func() {
			bodyBytes := reqTest("GET", fmt.Sprintf("/game/%d/play/start", gameID), nil, 200, "", true)
			sentence := models.Sentence{}
			err := json.Unmarshal(bodyBytes, &sentence)
			So(err, ShouldBeNil)
			So(sentence.Id, ShouldBeGreaterThan, 0)
		})
		Convey("문제 풀기 답 아님", func() {
			bodyBytes := reqTest("POST", fmt.Sprintf("/game/%d/play", gameID), models.Play{
				SentenceId: sentenceID,
				Input:      "xxxx",
			}, 201, "", true)
			play := models.Play{}
			err := json.Unmarshal(bodyBytes, &play)
			So(err, ShouldBeNil)
			So(play.IsSuccess, ShouldEqual, false)
		})
		Convey("문제 풀기 답 맞음", func() {
			bodyBytes := reqTest("POST", fmt.Sprintf("/game/%d/play", gameID), models.Play{
				SentenceId: sentenceID,
				Input:      "I study",
			}, 201, "", true)
			play := models.Play{}
			err := json.Unmarshal(bodyBytes, &play)
			So(err, ShouldBeNil)
			So(play.IsSuccess, ShouldEqual, true)
		})
		Convey("문제가 없어 문제 받기 실패", func() {
			reqTest("GET", fmt.Sprintf("/game/%d/play/start", gameID), nil, 500, "", true)
		})
		Convey("문제 풀기 끝", func() {
			bodyBytes := reqTest("PUT", fmt.Sprintf("/game/%d/end", gameID), nil, 200, "", true)
			game := models.Game{}
			err := json.Unmarshal(bodyBytes, &game)
			So(err, ShouldBeNil)
			So(game.Ended, ShouldEqual, true)
			So(game.Point, ShouldEqual, 1)
		})
		Convey("게임 종료로 문제 받기 실패", func() {
			reqTest("GET", fmt.Sprintf("/game/%d/play/start", gameID), nil, 400, "", true)
		})
	})
}

func TestGamePlay2(t *testing.T) {

	reqTest("POST", "/sentence", models.Sentence{
		Text:        "I study3",
		Korean:      "나는 공부한다3",
		TensesType:  tenses.Past,
		VerbsType:   verbs.General,
		PersonsType: persons.I,
		FormatsType: formats.Plain,
	}, 201, "", false)

	reqTest("POST", "/sentence", models.Sentence{
		Text:        "I study4",
		Korean:      "나는 공부한다4",
		TensesType:  tenses.Past,
		VerbsType:   verbs.General,
		PersonsType: persons.I,
		FormatsType: formats.Plain,
	}, 201, "", false)

	game := models.Game{
		TensesTypes:  "past",
		FormatsTypes: "",
		VerbsTypes:   "",
		PersonsTypes: "",
	}
	gameBodyBytes := reqTest("POST", "/game", &game, 201, "", false)
	json.Unmarshal(gameBodyBytes, &game)
	gameID := game.Id

	Convey("문제 받고 풀기", t, func() {
		Convey("문제 받고 풀기 2번", func() {
			bodyBytes := reqTest("GET", fmt.Sprintf("/game/%d/play/start", gameID), nil, 200, "", true)
			sentence := models.Sentence{}
			err := json.Unmarshal(bodyBytes, &sentence)
			So(err, ShouldBeNil)
			So(sentence.Id, ShouldBeGreaterThan, 0)

			bodyBytes = reqTest("POST", fmt.Sprintf("/game/%d/play", gameID), models.Play{
				SentenceId: sentence.Id,
				Input:      sentence.Text,
			}, 201, "", true)
			play := models.Play{}
			err = json.Unmarshal(bodyBytes, &play)
			So(err, ShouldBeNil)
			So(play.IsSuccess, ShouldEqual, true)

			bodyBytes = reqTest("GET", fmt.Sprintf("/game/%d/play/start", gameID), nil, 200, "", true)
			sentence = models.Sentence{}
			err = json.Unmarshal(bodyBytes, &sentence)
			So(err, ShouldBeNil)
			So(sentence.Id, ShouldBeGreaterThan, 0)

			bodyBytes = reqTest("POST", fmt.Sprintf("/game/%d/play", gameID), models.Play{
				SentenceId: sentence.Id,
				Input:      sentence.Text,
			}, 201, "", true)
			play = models.Play{}
			err = json.Unmarshal(bodyBytes, &play)
			So(err, ShouldBeNil)
			So(play.IsSuccess, ShouldEqual, true)

			bodyBytes = reqTest("PUT", fmt.Sprintf("/game/%d/end", gameID), nil, 200, "", true)
			game := models.Game{}
			err = json.Unmarshal(bodyBytes, &game)
			So(err, ShouldBeNil)
			So(game.Ended, ShouldEqual, true)
			So(game.Point, ShouldEqual, 2)
		})
	})
}
