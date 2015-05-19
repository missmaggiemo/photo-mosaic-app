
package controllers

import (
	"github.com/astaxie/beego"
//	"io/ioutil"
    "fmt"
    "photomosaic/imgproc"
//    "reflect"
)

type ProcessController struct {
	beego.Controller
}

func (c *ProcessController) Post() {
    f, hdr, _ := c.Ctx.Request.FormFile("file")
    defer f.Close()

    tile := imgproc.LoadImageFromStream(hdr.Filename, f)
    imgproc.SaveImage("tmp/" + hdr.Filename, tile)
    c.Ctx.WriteString(hdr.Filename)

    fmt.Println("Finished")
}
