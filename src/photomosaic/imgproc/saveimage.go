package imgproc


import "os"
import "image"
import "image/jpeg"


func SaveImage(path string, img image.Image) {
    toimg, _ := os.Create(path)
    defer toimg.Close()

    jpeg.Encode(toimg, img, &jpeg.Options{jpeg.DefaultQuality})
}