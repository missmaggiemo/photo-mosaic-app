package main

import (
	"fmt"
    "image"
	"photomosaic/tests"
	"photomosaic/imgproc"
)

func main() {
	var tile_paths []string = tests.GetTilePaths()
    var main_img_path string = tests.GetMainFilePath()

    var main_image = imgproc.LoadImage(main_img_path)

    main_image = imgproc.EnlargeMainImage(main_image)

    var tiles = make([]image.Image, len(tile_paths))

    for idx, tile_path := range tile_paths {
        fmt.Print(".")
        //fmt.Println(tile_path)
        tiles[idx] = imgproc.LoadImage(tile_path)
        tiles[idx] = imgproc.ResizeTile(tiles[idx])
    }

    var result_image image.Image = imgproc.Compose(main_image, tiles)

    imgproc.SaveImage("new.jpg", result_image)

}
