package tests

import "strings"
import "io/ioutil"
import "path/filepath"

func GetTilePaths(path string) []string {
    if path == "" {
        path = "./tests/sample_images/tiles"
    }

    file_info, _ := ioutil.ReadDir(path)

    var paths = make([]string, len(file_info))

    var idx int = 0
    for _, item := range file_info {
        if strings.Contains(item.Name(), ".png") || strings.Contains(item.Name(), ".jpg") {
            paths[idx], _ = filepath.Abs(path + item.Name())
            idx++
        }
    }

    return paths[0:idx]
}

func GetMainFilePath(path string) string {
    if path == "" {
        path = "/Users/psk/tubular/gohack/photo-mosaic-app/src/photomosaic/tests/sample_images/SunLou2.jpg"
    }

    res, _ := filepath.Abs(path)
    return res
}