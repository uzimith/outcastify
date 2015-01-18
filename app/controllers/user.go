package controllers

import (
	"github.com/k0kubun/pp"
	"github.com/revel/revel"
	"github.com/uzimith/outcastify/app/models"
	"github.com/uzimith/outcastify/app/routes"
)

type User struct {
	*revel.Controller
}

func (c User) Add(name string, id string) revel.Result {
	revel.INFO.Printf("user add:%s", name)
	// token := helper.GenerateRandom(16)
	// c.Session["token"] = token
	user := models.User{Name: name, Room: id, Join: true} //, Token: token}
	Gdb.Save(user)
	return c.Redirect(routes.App.Room(id))
}

func (c User) List(room string) revel.Result {
	var users []models.User
	Gdb.Where("Room = ?", room).Find(&users)
	return c.RenderJson(users)
}

func (c User) Share() revel.Result {
	var join map[int64]string
	var private map[int64]string
	var public string
	var group string

	c.Params.Bind(&join, "join")
	c.Params.Bind(&private, "private")
	c.Params.Bind(&public, "public")
	c.Params.Bind(&group, "group")

	pp.Println(join)
	pp.Println(private)
	pp.Println(public)
	pp.Println(group)

	return c.Redirect(routes.App.Room(c.Session["room"]))
}

func (c User) Delete(id int64) revel.Result {
	var user models.User
	Gdb.First(&user, id)
	Gdb.Delete(&user)
	return c.Redirect(App.Index)
}
