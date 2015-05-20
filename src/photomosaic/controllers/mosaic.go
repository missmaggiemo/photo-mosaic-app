package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"fmt"
	"photomosaic/imgproc"
)

type MosaicController struct {
	beego.Controller
}

func (c *MosaicController) Get() {
    if _, err := os.Stat("tmp/target.jpg"); err == nil {
        imgproc.Mosaic("tmp/target.jpg", "tmp/tiles")
        c.Ctx.WriteString("result.jpg")
    } else {
        imgproc.Mosaic("tmp/target.png", "tmp/tiles")
        c.Ctx.WriteString("result.jpg")
    }

    fmt.Println("Finished")
}
