package model

import (
	"ccs/db"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Student struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	CId int64 `json:"cid"`
}

func (s *Student) UserIsExists() bool {
	sql := "select id from students where id=?"
	var id int64
	_ = db.DB.QueryRow(sql, s.Id).Scan(&id)
	if id > 0 {
		return true
	}
	return false
}

func (s *Student) Add () {
	sql := "insert into students values(?,?,?,?)"
	_, err := db.DB.Exec(sql, s.Id, s.Name, s.Password, s.CId)
	if err != nil {
		log.Println(err.Error())
	}
}

func (s *Student) CheckPassword() bool {
	sql := "select password from student where id=?"
	var password string
	_ = db.DB.QueryRow(sql, s.Id).Scan(&password)
	if s.Password == password {
		return true
	}
	return false
}

