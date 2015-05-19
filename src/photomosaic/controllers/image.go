package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"io/ioutil"
	"fmt"
)

type ImageController struct {
	beego.Controller
}

func (c *ImageController) Get() {

    f, _ := os.Open("tmp/" + c.GetString("file"))
    defer f.Close()

    c.Ctx.Output.Header("Content-Type", "image/jpeg")

    res, _ := ioutil.ReadAll(f)
    c.Ctx.Output.Body(res)

    fmt.Println("Finished")
}
