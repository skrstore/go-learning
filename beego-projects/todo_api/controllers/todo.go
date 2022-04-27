package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"

	"todo_api/models"
)

type TodoController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Todo	true		"The todo content"
// @Success 200 {string} models.Todo.Id
// @Failure 403 body is empty
// @router / [post]
func (t *TodoController) Post() {
	var data models.Todo
	json.Unmarshal(t.Ctx.Input.RequestBody, &data)

	o := orm.NewOrm()

	todo := new(models.Todo)
	todo.Title = data.Title
	todo.Description = data.Description

	result, err := o.Insert(todo)
	if err != nil {
		fmt.Println("'Error occurred in creating Todo")
	} else {
		fmt.Println("Result ", result)
	}

	t.Data["json"] = map[string]string{"msg": "todo API"}
	t.ServeJSON()

	// // user := new(User)
	// // user.Profile = profile
	// // user.Name = "Admin"

	// // // // Insert
	// // fmt.Println(o.Insert(profile))
	// // fmt.Println(o.Insert(user))

	// // Read
	// // user1 := User{Id: 1}
	// // err = o.Read(&user1)
	// // fmt.Println("User", user1)
	// // fmt.Println("User E", err)

	// qs := o.QueryTable("user")

	// var users []*User
	// result, err := qs.All(&users)

	// fmt.Println("Res", result)
	// fmt.Println("Err", err)
	// fmt.Println("Users", users[0].Profile)

	// profileQS := o.QueryTable("profile")

	// var profiles []*Profile
	// result1, err1 := profileQS.All(&profiles)

	// fmt.Println("Result 1", result1)
	// fmt.Println("Error ", err1)
	// fmt.Println("AA", profiles[0])
}
