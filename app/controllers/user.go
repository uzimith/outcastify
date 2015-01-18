package controllers

import (
	"github.com/revel/revel"
	"github.com/uzimith/outcastify/app/models"
	"github.com/uzimith/outcastify/app/routes"
)

type User struct {
	*revel.Controller
}

func (c User) Add(name string, room string) revel.Result {
	revel.INFO.Printf("User.Add:%s %s", room, name)
	Gdb.Create(&models.User{Name: name, Room: room, Join: true})
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

	for id, _ := range private {
		Gdb.Where(models.Secret{
			Users: []models.User{{Id: id}},
			Room:  c.Session["room"]},
		).Assign(models.Secret{Text: "text"}).FirstOrCreate(models.Secret{})
	}
	Gdb.Create(&models.Secret{
		Users: []models.User{},
		Room:  c.Session["room"],
		Text:  public,
	})

	return c.Redirect(routes.App.Room(c.Session["room"]))
}

func (c User) Delete(id int64) revel.Result {
	var user models.User
	Gdb.First(&user, id)
	Gdb.Delete(&user)
	return c.Redirect(App.Index)
}
