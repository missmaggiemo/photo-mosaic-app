package imgproc

import (
    "image"
    "math"
    "image/draw"
//    "fmt"
)

func getHist(img image.Image) [100]int {
    var h [100]int

    var bounds image.Rectangle = img.Bounds()

    var sx, sy, fx, fy = bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y

    max_val := 65536 * 65536 * 3
    for x := sx; x < fx; x += 1 {
        for y := sy; y < fy; y += 1 {
            var ar, ag, ab, _ = img.At(x, y).RGBA()
            h[int(float64(ar*ar + ag*ag + ab*ab) / float64(max_val) * 99)] += 1
        }
    }

    return h
}

func diff3(ha [100]int, hb [100]int, min_diff float64) float64 {
    var res float64 = 0.
    for i:=0; i < 100; i++ {
        if res > min_diff {  // optimizitaion: obviously should stop
            return res
        }
        res += float64((ha[i] - hb[i]) * (ha[i] - hb[i]))
    }

    return res
}

func getTileWithMinimalDiff(part image.Image, tiles []image.Image, tile_hists [][100]int) image.Image {
    var min_diff float64 = 1e15
    var min_idx int = 0

    part_hist := getHist(part)

    for idx, _ := range tiles {
        var curr_diff = diff3(part_hist, tile_hists[idx], min_diff)
        if min_diff > curr_diff {
            min_diff = curr_diff
            min_idx = idx
        }
    }

    return tiles[min_idx]
}


type empty struct {}


func Compose(main_image image.Image, tiles []image.Image) image.Image {
    var bounds image.Rectangle = main_image.Bounds()
    var sx, sy, fx, fy = bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y

    var resulting_image = image.NewRGBA(bounds)

    var tile_hists = make([][100]int, len(tiles))
    for idx, tile := range tiles {
        tile_hists[idx] = getHist(tile)
    }

    N := int(math.Floor(float64(fx - sx) / TILE_WIDTH))
    M := int(math.Floor(float64(fy - sy) / TILE_HEIGHT))

    sem := make(chan empty, N*M)

    for x := 0; x < N; x++ {
        for y := 0; y < M; y++ {

            go func (x int, y int) {
                var rect = image.Rect(x * TILE_WIDTH, y * TILE_HEIGHT, (x + 1) * TILE_WIDTH, (y + 1) * TILE_HEIGHT)

                var subimg = main_image.(interface {
                    SubImage(r image.Rectangle) image.Image
                }).SubImage(rect)

                var tile = getTileWithMinimalDiff(subimg, tiles, tile_hists)
                draw.Draw(resulting_image, rect, tile, image.Point{0, 0}, draw.Src)

                sem <- empty{}
            } (x, y)

        }
    }

    proceed := 0.
    for i:= 0; i < N * M; i++ {
        <-sem
        proceed += 1
//        fmt.Print(int(proceed / float64(N * M) * 100.), "%, ")
    }

    return resulting_image
}