package controllers

import (
	"github.com/revel/revel"
	"github.com/uzimith/outcastify/app/helper"
	"github.com/uzimith/outcastify/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	id := helper.GenerateRandom(16)
	return c.Render(id)
}

func (c App) Room(id string) revel.Result {
	var users []models.User
	Gdb.Where("Room = ?", id).Find(&users)
	return c.Render(id, users)
}
