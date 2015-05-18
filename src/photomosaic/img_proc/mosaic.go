package img_proc

import (
    "github.com/nfnt/resize"
    "io/ioutil"
    "image/jpeg"
    "image/png"
    "os"
    "fmt"
)

const TILE_WIDTH = 50
const TILE_HEIGHT = 50


func resizeTiles(tiles &[]image.Image): {
    for idx, image := range *tiles {
        tiles[idx] := resize.Resize(TILE_WIDTH, TILE_HEIGHT, image, resize.Lanczos3)
    }
}


func Compose(main_image image.Image, tiles []image.Image) {

}

func main() {

    f, err := os.Open("test.jpg")
    data := jpeg.Decode(f)
    f.Close()

    var tiles_paths []string{"../"}
    var tiles [5]image.Image

    for i := 0; i < 5; i++ {
        f, err := os.Open( fmt.Sprintf("tiles/{}.jpg", string(i)) )
        jpeg.Decode(f)
        f.Close()
    }

}
