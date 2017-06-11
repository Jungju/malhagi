package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/jungju/malhagi/models"

	"time"

	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// PlayController operations for Play
type PlayController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description create Play
// @Param	body		body 	models.Play	true		"body for Play content"
// @Success 201 {int} models.Play
// @Failure 403 body is empty
// @router / [post]
func (c *PlayController) Post() {
	var v models.Play
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.CustomAbort(400, "Json 포멧이 아님")
	}

	if !v.ValidCreate() {
		c.CustomAbort(400, "누락된 정보가 있습니다.")
	}

	gameIDStr := c.Ctx.Input.Param(":gameID")
	gameID, _ := strconv.ParseInt(gameIDStr, 0, 64)
	game, err := models.GetGameById(gameID)
	if err != nil {
		c.CustomAbort(400, "Game 정보가 잘못되었습니다.")
	}

	if game.Ended {
		c.CustomAbort(400, "Game이 이미 끝났습니다.")
	}

	if game.CreatedAt.Add(time.Minute).After(time.Now()) {
		c.CustomAbort(400, "Game이 이미 끝났습니다.")
	}

	sentence, err := models.GetSentenceById(v.SentenceId)
	if err != nil {
		c.CustomAbort(400, "문제가 잘못되었습니다.")
	}

	v.GameId = game.Id
	if strings.ToLower(v.Input) == strings.ToLower(sentence.Text) {
		v.IsSuccess = true
	} else {
		v.IsSuccess = false
	}

	if _, err := models.AddPlay(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.CustomAbort(500, "Error System")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title 문제 가져오기
// @Description get Play by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Play
// @router /start [get]
func (c *PlayController) GetOne() {
	gameIDStr := c.Ctx.Input.Param(":gameID")
	gameID, _ := strconv.ParseInt(gameIDStr, 0, 64)
	game, err := models.GetGameById(gameID)
	if err != nil {
		c.CustomAbort(400, "Game 정보가 잘못되었습니다.")
	}

	if game.Ended {
		c.CustomAbort(400, "Game이 이미 끝났습니다.")
	}

	sentence, err := models.GetRandomSentence(game)
	if err != nil {
		c.CustomAbort(500, "준비된 문제가 없습니다.")
	}

	//sentence.Text = ""
	c.Data["json"] = sentence
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @router / [get]
func (c *PlayController) GetAll() {
	l, err := models.GetAllSentence()
	if err != nil {
		c.CustomAbort(500, "Error System")
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Play
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Play	true		"body for Play content"
// @Success 200 {object} models.Play
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PlayController) Put() {
	v := models.Play{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.CustomAbort(400, "Json이 잘못되었습니다.")
	}

	if !v.ValidUpdate() {
		c.CustomAbort(400, "누락된 정보가 있습니다.")
	}

	_, err := models.GetPlayById(v.Id)
	if err != nil {
		if err == orm.ErrNoRows {
			c.CustomAbort(404, "문장이 없습니다.")
		}
		c.CustomAbort(500, "Error System")
	}

	if err := models.UpdatePlay(&v); err != nil {
		c.CustomAbort(500, "Error System")
	}
	c.Ctx.Output.SetStatus(204)
}
