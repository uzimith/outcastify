package controllers

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"github.com/uzimith/outcastify/app/models"
)

// it can be used for jobs
var Gdb gorm.DB

// init db
func InitDB() {
	var err error
	// open db
	if revel.RunMode == "dev" {
		Gdb, err = gorm.Open("sqlite3", "/tmp/gorm.db")
		Gdb.LogMode(true)
	} else {
		Gdb, err = gorm.Open("postgres", "user=uname dbname=udbname sslmode=disable password=supersecret")
	}
	if err != nil {
		revel.ERROR.Println("database error:", err)
		panic(err)
	}
	revel.INFO.Println("DB:migration!")
	Gdb.AutoMigrate(&models.User{})
}

// transactions
type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

// This method fills the c.Txn before each transaction
func (c *GormController) Begin() revel.Result {
	txn := Gdb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	c.Txn = txn
	return nil
}

// This method clears the c.Txn after each transaction
func (c *GormController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Commit()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

// This method clears the c.Txn after each transaction, too
func (c *GormController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Rollback()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
