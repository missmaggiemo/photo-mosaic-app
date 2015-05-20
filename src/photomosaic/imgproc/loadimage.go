package imgproc

import (
    "os"
    "strings"
    "image"
    "image/jpeg"
    "image/png"
    "mime/multipart"
    "fmt"
)

func LoadImage(path string) image.Image {
    f, err := os.Open(path)

    defer f.Close()

    var data image.Image

    if strings.Contains(path, ".png") {
        data, err = png.Decode(f)

        //fmt.Println("HERE ERROR: ", err, path)
        if err != nil {
            data, err = jpeg.Decode(f)
            //fmt.Println("JPG NOT HELPED: ", err, path)
        }
    } else {
        data, _ = jpeg.Decode(f)
    }
    return data
}


func LoadImageFromStream(path string, f multipart.File) image.Image {
    if strings.Contains(path, ".png") {
        data, err := png.Decode(f)
        fmt.Println("ERROR: ", err, path)
        if err == nil {
            return data
        } else {
            return nil
        }
    } else {
        data, err := jpeg.Decode(f)
        fmt.Println("ERROR: ", err, path)
        if err == nil {
            return data
        } else {
            return nil
        }
    }
}