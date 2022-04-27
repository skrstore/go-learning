package main

import (
	"github.com/beego/beego/v2/server/web"
)

type MainController struct {
	web.Controller
}

func (m *MainController) Get() {
	m.Data["userName"] = "SuperAdmin"
	m.TplName = "index.tpl"
}

type User struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age,text,age:"`
	Sex   string
	Intro string `form:",textarea"`
}

type FormRenderController struct {
	web.Controller
}

func (f *FormRenderController) Get() {
	f.Data["Form"] = &User{}
	f.TplName = "form_render.tpl"
}

func main() {
	web.Router("/", &MainController{})
	web.Router("/form_render", &FormRenderController{})
	web.Run()
}
