package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"photomosaic/imgproc"
	"strings"
	"io/ioutil"
	"path/filepath"
)

type MosaicController struct {
	beego.Controller
}

func (c *MosaicController) Get() {

    file_info, _ := ioutil.ReadDir("tmp")

    var main_file string

    for _, item := range file_info {
        if (strings.Contains(item.Name(), ".png") || strings.Contains(item.Name(), ".jpg")) && (item.Name() != "result.jpg") {
            main_file, _ = filepath.Abs("tmp/" + item.Name())
        }
    }

    imgproc.Mosaic(main_file, "tmp/tiles/")
    c.Ctx.WriteString("result.jpg")

    fmt.Println("Finished")
}
