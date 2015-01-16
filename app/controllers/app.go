package controllers

import (
	"github.com/revel/revel"
	"github.com/uzimith/outcastify/app/room"
)

import "code.google.com/p/go.net/websocket"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Socket(user string, ws *websocket.Conn) revel.Result {
	revel.INFO.Printf("new connection:" + user)
	revel.INFO.Printf("%+v", ws)
	event := room.NewEvent("join", "me", "I got")
	if websocket.JSON.Send(ws, event) != nil {
		return nil
	}
	newMessages := make(chan string)
	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	for {
		select {
		case msg, ok := <-newMessages:
			if !ok {
				return nil
			}

			revel.INFO.Println(msg)
			// room.Say(user, msg)
		}
	}
	return nil
}
