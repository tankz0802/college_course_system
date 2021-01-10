package model

import (
	"ccs/db"
	"log"
)

type Class struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	TId  int64  `json:"tid"`
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

func GetClassList() ([]*Class, error) {
	sql := "select id,name from class"
	rows, err := db.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	classList := make([]*Class, 0)
	for rows.Next() {
		class := &Class{}
		if err := rows.Scan(&class.Id, &class.Name); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		classList = append(classList, class)
	}
	return classList, nil
}

type ClassInfo struct {
	Name    string `json:"name"`
	Teacher string `json:"teacher"`
}

func GetClassInfoList() ([]*ClassInfo, error) {
	sql := "select c.name,t.name from class as c inner join teacher as t on c.tid=t.id"
	rows, err := db.DB.Query(sql)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	classInfoList := make([]*ClassInfo, 0)
	for rows.Next() {
		classInfo := &ClassInfo{}
		if err := rows.Scan(&classInfo.Name, &classInfo.Teacher); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		classInfoList = append(classInfoList, classInfo)
	}
	return classInfoList, nil
}