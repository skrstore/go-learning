package main

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

// Query parameter
type QueryParamController struct {
	web.Controller
}

func (q *QueryParamController) Get() {
	// URL: http://127.0.0.1:8080/?name=Admin
	name := q.GetString("name")
	fmt.Println("Name ", name)

	q.Ctx.WriteString("hello world")
}

// Parameterized URL
type ParametrizedController struct {
	web.Controller
}

func (p *ParametrizedController) Get() {
	// URL: http://127.0.0.1:8080/1001
	id := p.Ctx.Input.Param(":id")
	fmt.Println("User ID ", id)

	p.Ctx.WriteString("Parametrized URL Check \nVisit: http://127.0.0.1:8080/1001")
}

// Getting Data from GET Form
type GetFormController struct {
	web.Controller
}

func (g *GetFormController) Get() {
	// URL : http://127.0.0.1:8080/sendData
	category := g.GetString("category")
	color := g.GetString("color")

	fmt.Println("category ", category)
	fmt.Println("color ", color)
	g.Ctx.WriteString("Get Data \nCategory : " + category + " \nColor : " + color)
}

// Getting Data from POST Form
type PostFormController struct {
	web.Controller
}

func (g *PostFormController) Post() {
	// URL : http://127.0.0.1:8080/sendData
	category := g.GetString("category")
	color := g.GetString("color")

	fmt.Println("category ", category)
	fmt.Println("color ", color)
	g.Ctx.WriteString("Post Data \nCategory : " + category + " \nColor : " + color)
}

func main() {
	web.Router("/", &QueryParamController{})
	web.Router("/:id", &ParametrizedController{})
	web.Router("/send_data_get", &GetFormController{})
	web.Router("/send_data_post", &PostFormController{})

	web.SetStaticPath("/sendData", "public/index.html")

	web.Run()
}
