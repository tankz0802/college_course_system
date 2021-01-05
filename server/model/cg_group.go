package model

import (
	"ccs/db"
	"log"
)

type CgGroup struct {
	CgId int64 `json:"cg_id"`
	TId int64 `json:"tid"`
}

func (cg *CgGroup) Add() error {
	sql := "insert into cg_group values(?,?)"
	_, err := db.DB.Exec(sql, cg.CgId, cg.TId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}