package controllers

import (
	"strconv"

	"github.com/jungju/malhagi/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GameController operations for Game
type GameController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description create Game
// @Param	body		body 	models.Game	true		"body for Game content"
// @Success 201 {int} models.Game
// @Failure 403 body is empty
// @router / [post]
func (c *GameController) Post() {
	game := &models.Game{
		Ended: false,
		Point: 0,
	}
	if _, err := models.AddGame(game); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = game
	} else {
		c.CustomAbort(500, "Error System")
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
	v, err := models.GetGameById(id)
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
// @Description update the Game
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Game	true		"body for Game content"
// @Success 200 {object} models.Game
// @Failure 403 :id is not int
// @router /:id/end [put]
func (c *GameController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	game, err := models.GetGameById(id)
	if err != nil {
		if err == orm.ErrNoRows {
			c.CustomAbort(404, "게임이 없습니다.")
		}
		c.CustomAbort(500, "Error System")
	}

	game.Ended = true
	if cnt, err := models.GetPlaySuccessCountByGameID(id); err == nil {
		game.Point = cnt
	} else {
		c.CustomAbort(500, "Error System")
	}

	if err := models.UpdateGame(game); err != nil {
		c.CustomAbort(500, "Error System")
	}
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = game
	c.ServeJSON()
}
