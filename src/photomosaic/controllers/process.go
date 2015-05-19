
package controllers

import (
	"github.com/astaxie/beego"
    "fmt"
    "reflect"
)

type ProcessController struct {
	beego.Controller
}

func (c *ProcessController) Post() {
    f, h, _ := c.GetFile("tiles")

	fmt.Println(h)
	fmt.Println(reflect.TypeOf(f))

	c.TplNames = "index.html"
}
