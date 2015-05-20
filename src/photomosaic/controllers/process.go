
package controllers

import (
	"github.com/astaxie/beego"
//	"io/ioutil"
    "fmt"
    "photomosaic/imgproc"
//    "reflect"
    "os"
    "strings"
    "io/ioutil"
    "path/filepath"
)

type ProcessController struct {
	beego.Controller
}

func (c *ProcessController) Post() {
    f, hdr, err := c.Ctx.Request.FormFile("target")
    is_tile := false

    os.MkdirAll("tmp/tiles", 0777)

    if err != nil {
        is_tile = true
        f, hdr, _ = c.Ctx.Request.FormFile("tile")
    } else {

        file_info, _ := ioutil.ReadDir("tmp")

        var main_file string

        for _, item := range file_info {
            if (strings.Contains(item.Name(), ".png") || strings.Contains(item.Name(), ".jpg")) && (item.Name() != "result.jpg") {
                main_file, _ = filepath.Abs("tmp/" + item.Name())
                err := os.Remove(main_file)
                if err != nil {
                    fmt.Println("Oops")
                }
            }
        }

    }

    defer f.Close()

    tile := imgproc.LoadImageFromStream(hdr.Filename, f)

    if tile != nil {

        if is_tile {
            tile = imgproc.ResizeTile(tile)
            imgproc.SaveImage("tmp/tiles/" + hdr.Filename, tile)
            c.Ctx.WriteString("tiles/" + hdr.Filename)
        } else {
            imgproc.SaveImage("tmp/" + hdr.Filename, tile)
            c.Ctx.WriteString(hdr.Filename)
        }

    }

    c.Ctx.WriteString("nil")
    fmt.Println("Finished")
}
