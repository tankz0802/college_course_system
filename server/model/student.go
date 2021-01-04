package model

import _ "github.com/go-sql-driver/mysql"

type Student struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	CId int64 `json:"cid"`
}

func (s *Student) UserIsExists() bool {
	sql := "select id from students where id=?"

}

func (s *Student) Add () {
	sql := ""
}