package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/jungju/malhagi/models"

	"github.com/astaxie/beego"
)

// GameController operations for Game
type GameController struct {
	beego.Controller
}

// URLMapping ...
func (c *GameController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
}

// Post ...
// @Title Post
// @Description create Game
// @Param	body		body 	models.Game	true		"body for Game content"
// @Success 201 {int} models.Game
// @Failure 403 body is empty
// @router / [post]
func (c *GameController) Post() {
	var v models.Game
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.AddGame(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Game by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Game
// @Failure 403 :id is empty
// @router /:id [get]
func (c *GameController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetGameById(int(id))
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @router / [get]
func (c *GameController) GetAll() {
	l, err := models.GetAllGame()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Game
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Game	true		"body for Game content"
// @Success 200 {object} models.Game
// @Failure 403 :id is not int
// @router /:id [put]
func (c *GameController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Game{Id: int(id)}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateGameById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
