package controllers

import (
	"github.com/revel/revel"
	"github.com/uzimith/outcastify/app/helper"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	id := helper.GenerateRandom(16)
	return c.Render(id)
}

func (c App) Room(room string) revel.Result {
	c.Session["room"] = room
	userId := c.Session["userId"]
	token := c.Session["token"]
	return c.Render(room, userId, token)
}
