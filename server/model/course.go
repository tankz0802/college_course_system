package model

import (
	"ccs/db"
	"log"
)

type Course struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Duration int32 `json:"duration"`
	Credit int32 `json:"credit"`
	Category string `json:"category"`
	Num int32 `json:"num"`
}

func (c *Course) CourseIsExists() bool {
	sql := "select id from course where id=?"
	var id int64
	db.DB.QueryRow(sql, c.Id).Scan(&id)
	if id > 0 {
		return true
	}
	return false
}

func (c *Course) Add() error {
	sql := "insert into course values(?,?,?,?,?,?)"
	_, err := db.DB.Exec(sql, c.Id, c.Name, c.Duration, c.Credit, c.Category, c.Num)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
