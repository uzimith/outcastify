package controllers

import (
	"github.com/revel/revel"
	"github.com/uzimith/outcastify/app/helper"
	"github.com/uzimith/outcastify/app/models"
	"github.com/uzimith/outcastify/app/routes"
)

type User struct {
	*revel.Controller
}

func (c User) Add(name string, id string) revel.Result {
	revel.INFO.Printf("user add:%s", name)
	token := helper.GenerateRandom(16)
	c.Session["token"] = token
	user := models.User{Name: name, Room: id, Join: true, Token: token}
	Gdb.Save(user)
	return c.Redirect(routes.App.Room(id))
}

func (c User) Delete(id int64) revel.Result {
	var user models.User
	Gdb.First(&user, id)
	Gdb.Delete(&user)
	return c.Redirect(App.Index)
}
