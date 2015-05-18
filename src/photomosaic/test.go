package main

import (
	"photomosaic/tests"
	"fmt"
)

func main() {
	//var res []string = tests.GetTilePaths()
    var main_img string = tests.GetMainFilePath()

    fmt.Println(main_img)

	//for _, x := range res {
	//	fmt.Println(x)
	//}
}



