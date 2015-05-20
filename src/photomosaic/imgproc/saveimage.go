package imgproc


import "os"
import "image"
import "image/jpeg"
import "image/png"
import "strings"


func SaveImage(path string, img image.Image) {
    toimg, _ := os.Create(path)
    defer toimg.Close()

    if strings.Contains(path, ".png") {
        png.Encode(toimg, img)
    } else {
        jpeg.Encode(toimg, img, &jpeg.Options{jpeg.DefaultQuality})
    }
}