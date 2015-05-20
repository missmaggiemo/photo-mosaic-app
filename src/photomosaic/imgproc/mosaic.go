package imgproc

import (
	"fmt"
    "image"
	"photomosaic/tests"
	"runtime"
)

func Mosaic(main_image_path, tiles_path string) {
    fmt.Println("# of proc: ", runtime.NumCPU())
    runtime.GOMAXPROCS(runtime.NumCPU())

    //var tiles_path = ""
    //var main_image_path = ""

    //if len(os.Args) >= 2 {
    //    main_image_path = os.Args[1]
    //    tiles_path = os.Args[2]
    //}

	var tile_paths []string = tests.GetTilePaths(tiles_path)
    var main_img_path string = tests.GetMainFilePath(main_image_path)

    var main_image = LoadImage(main_img_path)
    main_image = EnlargeMainImage(main_image)

    var tiles = make([]image.Image, len(tile_paths))

    sem := make(chan empty, len(tile_paths))

    for idx, tile_path := range tile_paths {
        go func (idx int, tile_path string) {
            tiles[idx] = LoadImage(tile_path)
            tiles[idx] = ResizeTile(tiles[idx])
            sem <- empty{}
        }(idx, tile_path)
    }

    for i:=0; i<len(tile_paths); i++ {
        <-sem
        fmt.Print(".")
    }

    var result_image image.Image = Compose(main_image, tiles)

    SaveImage("tmp/result.jpg", result_image)

}
