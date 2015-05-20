
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
    f, hdr, err := c.Ctx.Request.FormFile("target")
    is_tile := false

    if err != nil {
        is_tile = true
        f, hdr, _ = c.Ctx.Request.FormFile("tile")

        fmt.Println("hdr.Filename")
    }
    defer f.Close()

    tile := imgproc.LoadImageFromStream(hdr.Filename, f)

    if is_tile {
        imgproc.SaveImage("tmp/tiles/" + hdr.Filename, tile)
        c.Ctx.WriteString("tiles/" + hdr.Filename)
    } else {
        imgproc.SaveImage("tmp/" + hdr.Filename, tile)
        c.Ctx.WriteString(hdr.Filename)
    }

    fmt.Println("Finished")
}
