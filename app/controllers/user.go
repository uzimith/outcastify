package controllers

import (
	"fmt"

	"github.com/revel/revel"
	"github.com/uzimith/outcastify/app/helper"
	"github.com/uzimith/outcastify/app/models"
	"github.com/uzimith/outcastify/app/routes"
)

type User struct {
	*revel.Controller
}

func (c User) Add(name string, room string) revel.Result {
	token := helper.GenerateRandom(16)
	user := models.User{Name: name, Room: room, Join: true, Token: token}
	Gdb.Create(&user)
	c.Session["token"] = token
	c.Session["userId"] = fmt.Sprintf("%d", user.Id)
	revel.INFO.Printf("User.Add:%s %s(%d)", room, name, user.Id)
	return c.Redirect(routes.App.Room(room))
}

func (c User) List(room string) revel.Result {
	revel.INFO.Printf("User.List")
	var users []models.User
	Gdb.Where(&models.User{Room: room}).Find(&users)
	return c.RenderJson(users)
}

func (c User) Share() revel.Result {
	revel.INFO.Printf("User.Share")
	var join map[int64]string
	var private map[int64]string
	var public string
	var group string

	c.Params.Bind(&join, "join")
	c.Params.Bind(&private, "private")
	c.Params.Bind(&public, "public")
	c.Params.Bind(&group, "group")

	for id, text := range private {
		var user models.User
		Gdb.Where(&models.User{Id: id}).Find(&user)
		secret := models.Secret{Text: text}
		Gdb.Create(&secret)
		Gdb.Model(&user).Association("Secrets").Append(secret)
	}
	// for id := range join {
	// 	Gdb.Create(&models.Secret{
	// 		Users: []models.User{{Id: id}},
	// 		Room:  room,
	// 		Text:  public,
	// 	})
	// }

	return c.Redirect(routes.App.Room(c.Session["room"]))
}

func (c User) Delete(id int64) revel.Result {
	var user models.User
	Gdb.First(&user, id)
	Gdb.Delete(&user)
	return c.Redirect(App.Index)
}
