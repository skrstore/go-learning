package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Todo struct {
	Id          int
	Title       string
	Description string
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Todo))
}
