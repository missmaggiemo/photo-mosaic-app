package tests

import "strings"
import "io/ioutil"
import "path/filepath"

func GetTilePaths() []string {
    file_info, _ := ioutil.ReadDir("./tests/sample_images/tiles")

    var paths = make([]string, len(file_info))

    var idx int = 0
    for _, item := range file_info {
        if strings.Contains(item.Name(), ".png") || strings.Contains(item.Name(), ".jpg") {
            paths[idx], _ = filepath.Abs("./tests/sample_images/tiles/" + item.Name())
            idx++
        }
    }

    return paths[0:idx]
}

func GetMainFilePath() string {
    res, _ := filepath.Abs("/Users/psk/tubular/gohack/photo-mosaic-app/src/photomosaic/tests/sample_images/dogecoin-300.jpg")
    return res
}