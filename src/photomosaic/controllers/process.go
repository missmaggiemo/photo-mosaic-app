
package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
    "fmt"
//    "reflect"
)

type ProcessController struct {
	beego.Controller
}

func (c *ProcessController) Post() {
    f, _, _ := c.Ctx.Request.FormFile("file")
    defer f.Close()

    tile, _ := ioutil.ReadAll(f)

    c.Ctx.Output.Header("Content-Type", "image/jpeg")
    c.Ctx.Output.Body(tile)

    fmt.Println("Finished")
}
