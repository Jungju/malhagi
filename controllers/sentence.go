package controllers

import (
	"encoding/json"

	"github.com/jungju/malhagi/models"

	"github.com/astaxie/beego"
)

// SentenceController operations for Sentence
type SentenceController struct {
	beego.Controller
}

// URLMapping ...
func (c *SentenceController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
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
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.AddSentence(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Sentence
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Sentence
// @Failure 403
// @router / [get]
func (c *SentenceController) GetAll() {
	l, err := models.GetAllSentence()
	if err != nil {
		c.Data["json"] = err.Error()
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
	idStr := c.Ctx.Input.Param(":id")
	v := models.Sentence{Verb: idStr}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateSentenceById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
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
	if err := models.DeleteSentence(idStr); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
