package imgproc

import (
    "github.com/nfnt/resize"
    "image"
)

const TILE_WIDTH = 30
const TILE_HEIGHT = 30
const ENLARGE = 8

func ResizeTile(tile image.Image) image.Image {
    return resize.Resize(TILE_WIDTH, TILE_HEIGHT, tile, resize.Lanczos3)
}

func EnlargeMainImage(main_image image.Image) image.Image {
    bounds := main_image.Bounds()
    return resize.Resize(uint((bounds.Max.X - bounds.Min.X) * ENLARGE),
                         uint((bounds.Max.Y - bounds.Min.Y) * ENLARGE),
                         main_image, resize.Lanczos3)
}
