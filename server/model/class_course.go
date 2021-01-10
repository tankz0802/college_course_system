package model

import (
	"database/sql"
	"log"
)

type ClassCourse struct {
	Id int64 `json:"id"`
	Cid string `json:"cid"`
}

func (cc *ClassCourse) Add(tx *sql.Tx) error {
	sql := "insert into class_course values(?,?)"
	_, err := tx.Exec(sql, cc.Id, cc.Cid)
	if err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return err
	}
	return nil
}

