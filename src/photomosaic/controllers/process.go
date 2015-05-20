
package controllers

import (
	"github.com/astaxie/beego"
//	"io/ioutil"
    "fmt"
    "photomosaic/imgproc"
//    "reflect"
    "strings"
)

type ProcessController struct {
	beego.Controller
}

func (c *ProcessController) Post() {
    f, hdr, _ := c.Ctx.Request.FormFile("tile")
    defer f.Close()

    tile := imgproc.LoadImageFromStream(hdr.Filename, f)

    if strings.Contains(hdr.Filename, "tile") {
        imgproc.SaveImage("tmp/tiles/" + hdr.Filename, tile)
        c.Ctx.WriteString("tiles/" + hdr.Filename)
    } else {
        imgproc.SaveImage("tmp/" + hdr.Filename, tile)
        c.Ctx.WriteString(hdr.Filename)
    }

    fmt.Println("Finished")
}
