package controllers

import (
	"fmt"
	"time"

	"code.google.com/p/go.net/websocket"

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

func (c User) List(room string, ws *websocket.Conn) revel.Result {
	ticker := time.NewTicker(time.Millisecond * 500)
	func() {
		revel.INFO.Printf("User.List: Start - %s", c.Session["userId"])
		var sentAt time.Time
		for {
			select {
			case <-ticker.C:
				var lastUser models.User
				Gdb.Where(&models.User{Room: room}).Order("updated_at desc").Find(&lastUser)
				if lastUser.UpdatedAt.Sub((sentAt)) > 0 {
					sentAt = lastUser.UpdatedAt
					var users []models.User
					Gdb.Where(&models.User{Room: room}).Find(&users)
					if websocket.JSON.Send(ws, &users) != nil {
						revel.WARN.Printf("User.List: Send Error!")
						return
					}
				}
			}
		}
	}()
	revel.INFO.Printf("User.List: End - %s", c.Session["userId"])
	return nil
}

func (c User) Delete(id int64) revel.Result {
	var user models.User
	Gdb.First(&user, id)
	Gdb.Delete(&user)
	return nil
}
