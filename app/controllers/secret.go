package controllers

import (
	"strconv"
	"time"

	"code.google.com/p/go.net/websocket"

	"github.com/revel/revel"
)
import "github.com/uzimith/outcastify/app/models"

type Secret struct {
	*revel.Controller
}

func (c Secret) List(ws *websocket.Conn) revel.Result {
	id, _ := strconv.ParseInt(c.Session["userId"], 10, 64)
	ticker := time.NewTicker(time.Millisecond * 500)
	var user models.User
	Gdb.Where(&models.User{Id: id}).Find(&user)
	func() {
		revel.INFO.Printf("Secret.List: Start - %s", c.Session["userId"])
		var sentAt time.Time
		for {
			select {
			case <-ticker.C:
				var lastSecret models.Secret
				Gdb.Order("updated_at desc").Model(&user).Association("Secrets").Find(&lastSecret)
				if lastSecret.UpdatedAt.Sub((sentAt)) > 0 {
					revel.INFO.Printf("Secret.List: Send - %s", c.Session["userId"])
					sentAt = lastSecret.UpdatedAt
					var secrets []models.Secret
					Gdb.Model(&user).Association("Secrets").Find(&secrets)
					if websocket.JSON.Send(ws, &secrets) != nil {
						revel.WARN.Printf("Secret.List: Send Error!")
						return
					}
				}
			}
		}
	}()
	revel.INFO.Printf("Secret.List: End - %s", c.Session["userId"])
	return nil
}

func (c Secret) Add() revel.Result {
	revel.INFO.Printf("Secret.Add")
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
		secret := models.Secret{Type: "private", Text: text}
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

	return nil
}
