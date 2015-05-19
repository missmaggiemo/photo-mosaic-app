package imgproc

import (
    "github.com/nfnt/resize"
    "image"
)

const TILE_WIDTH = 10
const TILE_HEIGHT = 10

func ResizeTile(tile image.Image) image.Image {
    return resize.Resize(TILE_WIDTH, TILE_HEIGHT, tile, resize.Lanczos3)
}
