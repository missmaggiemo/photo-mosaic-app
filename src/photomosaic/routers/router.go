package routers

import (
	"photomosaic/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/process", &controllers.ProcessController{})
    beego.Router("/image", &controllers.ImageController{})
}
