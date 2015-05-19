package imgproc

import (
    "os"
    "strings"
    "image"
    "image/jpeg"
    "image/png"
)

func LoadImage(path string) image.Image {
    f, _ := os.Open(path)
    defer f.Close()

    var data image.Image

    if strings.Contains(path, ".png") {
        data, _ = png.Decode(f)
    } else {
        data, _ = jpeg.Decode(f)
    }
    return data
}