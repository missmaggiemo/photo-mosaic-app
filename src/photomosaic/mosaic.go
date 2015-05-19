package main

import (
	"fmt"
    "image"
	"photomosaic/tests"
	"photomosaic/imgproc"
	"os"
)


func main() {
    var tiles_path = ""
    var main_image_path = ""

    if len(os.Args) >= 2 {
        main_image_path = os.Args[1]
        tiles_path = os.Args[2]
    }

	var tile_paths []string = tests.GetTilePaths(tiles_path)
    var main_img_path string = tests.GetMainFilePath(main_image_path)

    var main_image = imgproc.LoadImage(main_img_path)

    main_image = imgproc.EnlargeMainImage(main_image)

    var tiles = make([]image.Image, len(tile_paths))

    for idx, tile_path := range tile_paths {
        fmt.Print(".")
        tiles[idx] = imgproc.LoadImage(tile_path)
        tiles[idx] = imgproc.ResizeTile(tiles[idx])
    }

    var result_image image.Image = imgproc.Compose(main_image, tiles)

    imgproc.SaveImage("new.jpg", result_image)

}
