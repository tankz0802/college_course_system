package model

import (
	"ccs/db"
	"log"
)

type Class struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	TId int64 `json:"tid"`
}

func (c *Class) ClassIsExists() bool {
	sql := "select id from class where id=?"
	var id int64
	_ = db.DB.QueryRow(sql, c.Id).Scan(&id)
	if id > 0 {
		return true
	}
	return false
}

func (c *Class) Add() {
	sql := "insert into class values(?,?,?)"
	_, err := db.DB.Exec(sql, c.Id, c.Name, c.TId)
	if err != nil {
		log.Println(err.Error())
	}
}