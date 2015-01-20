package controllers

import (
	"strconv"

	"github.com/revel/revel"
)
import "github.com/uzimith/outcastify/app/models"

type Secret struct {
	*revel.Controller
}

func (c Secret) List() revel.Result {
	id, _ := strconv.ParseInt(c.Session["userId"], 10, 64)

	revel.INFO.Printf("Secret.List:")
	var user models.User
	Gdb.Where(&models.User{Id: id}).Find(&user)

	var secrets []models.Secret

	Gdb.Debug().Model(&user).Association("Secrets").Find(&secrets)

	return c.RenderJson(secrets)
}
