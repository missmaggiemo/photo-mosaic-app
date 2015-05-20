package imgproc

import (
    "image"
    "math"
    "image/draw"
//    "fmt"
)

const HIST_SIZE = 100

func getHist(img image.Image) ([HIST_SIZE]int, [HIST_SIZE]int, [HIST_SIZE]int) {
    var hr [HIST_SIZE]int
    var hg [HIST_SIZE]int
    var hb [HIST_SIZE]int

    var bounds image.Rectangle = img.Bounds()

    var sx, sy, fx, fy = bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y

    max_val := 65536 * 65536 * 3
    for x := sx; x < fx; x += 1 {
        for y := sy; y < fy; y += 1 {
            var ar, ag, ab, _ = img.At(x, y).RGBA()
            hr[int(float64(ar*ar) / float64(max_val) * (HIST_SIZE-1))] += 1
            hg[int(float64(ag*ag) / float64(max_val) * (HIST_SIZE-1))] += 1
            hb[int(float64(ab*ab) / float64(max_val) * (HIST_SIZE-1))] += 1
        }
    }

    return hr, hg, hb
}

func diff3(har, hag, hab [HIST_SIZE]int, hbr, hbg, hbb [HIST_SIZE]int, min_diff float64) float64 {
    var res float64 = 0.
    for i:=0; i < HIST_SIZE; i++ {
        if res > min_diff {  // optimizitaion: obviously should stop
            return res
        }
        res += float64((har[i] - hbr[i]) * (har[i] - hbr[i]))
        res += float64((hag[i] - hbg[i]) * (hag[i] - hbg[i]))
        res += float64((hab[i] - hbb[i]) * (hab[i] - hbb[i]))
    }

    return res
}

func getTileWithMinimalDiff(part image.Image, tiles []image.Image, thr, thg, thb [][HIST_SIZE]int) image.Image {
    var min_diff float64 = 1e15
    var min_idx int = 0

    part_hist_r, part_hist_g, part_hist_b := getHist(part)

    for idx, _ := range tiles {
        if tiles[idx] == nil {
            continue
        }
        var curr_diff = diff3(part_hist_r, part_hist_g, part_hist_b,
                              thr[idx], thg[idx], thb[idx], min_diff)
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

    var tile_hists_r = make([][HIST_SIZE]int, len(tiles))
    var tile_hists_g = make([][HIST_SIZE]int, len(tiles))
    var tile_hists_b = make([][HIST_SIZE]int, len(tiles))
    for idx, tile := range tiles {
        if tiles == nil {
            continue
        }
        tile_hists_r[idx], tile_hists_g[idx], tile_hists_b[idx] = getHist(tile)
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

                var tile = getTileWithMinimalDiff(subimg, tiles, tile_hists_r, tile_hists_g, tile_hists_b)
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