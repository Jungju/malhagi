package controllers

import (
	"encoding/json"

	"github.com/jungju/malhagi/models"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// SentenceController operations for Sentence
type SentenceController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description create Sentence
// @Param	body		body 	models.Sentence	true		"body for Sentence content"
// @Success 201 {int} models.Sentence
// @Failure 403 body is empty
// @router / [post]
func (c *SentenceController) Post() {
	var v models.Sentence
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.CustomAbort(400, "Json 포멧이 아님")
	}

	if !v.ValidCreate() {
		c.CustomAbort(400, "누락된 정보가 있습니다.")
	}

	_, err := models.GetSentenceByText(v.Text)
	if err != nil {
		if err != orm.ErrNoRows {
			c.CustomAbort(500, "문장을 가져오는중 알수없는 에러")
		}
	} else {
		c.CustomAbort(409, "이미 영어가 입력되어 있음")
	}

	if _, err := models.AddSentence(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.CustomAbort(500, "DB에 추가 중 알수없는 에러")
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Sentence
// @Success 200 {object} models.Sentence
// @Failure 403
// @router / [get]
func (c *SentenceController) GetAll() {
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
// @Description update the Sentence
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Sentence	true		"body for Sentence content"
// @Success 200 {object} models.Sentence
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SentenceController) Put() {
	v := models.Sentence{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		c.CustomAbort(400, "Json이 잘못되었습니다.")
	}

	if !v.ValidUpdate() {
		c.CustomAbort(400, "누락된 정보가 있습니다.")
	}

	_, err := models.GetSentenceById(v.Id)
	if err != nil {
		if err == orm.ErrNoRows {
			c.CustomAbort(404, "문장이 없습니다.")
		}
		c.CustomAbort(500, "Error System")
	}

	if err := models.UpdateSentence(&v); err != nil {
		c.CustomAbort(500, "Error System")
	}
	c.Ctx.Output.SetStatus(204)
}

// Delete ...
// @Title Delete
// @Description delete the Sentence
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SentenceController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.CustomAbort(400, "ID가 없습니다.")
	}

	_, err = models.GetSentenceById(id)
	if err != nil {
		if err == orm.ErrNoRows {
			c.CustomAbort(404, "문장이 없습니다.")
		}
		c.CustomAbort(500, "Error System")
	}

	if err := models.DeleteSentence(id); err != nil {
		c.CustomAbort(500, "Error System")
	}
	c.Ctx.Output.SetStatus(204)
}
