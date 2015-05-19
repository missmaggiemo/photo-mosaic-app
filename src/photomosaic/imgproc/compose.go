package imgproc

import (
    "image"
    "math"
    "image/draw"
)


func diff(a, b image.Image) int {
    var res float64 = 0
    var bounds image.Rectangle = a.Bounds()

    var sx, sy, fx, fy = bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y
    for x := sx; x < fx; x += 1 {
        for y := sy; y < fy; y += 1 {
            var ar, ag, ab, aa = a.At(x, y).RGBA()
            var br, bg, bb, ba = b.At(x, y).RGBA()
            res += math.Sqrt(math.Pow(float64(ar - br), 2) +
                             math.Pow(float64(ag - bg), 2) +
                             math.Pow(float64(ab - bb), 2) +
                             math.Pow(float64(aa - ba), 2))
        }
    }

    return int(res)
}

func diff2(a, b image.Image) int {
    var bounds image.Rectangle = a.Bounds()

    var sx, sy, fx, fy = bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y

    var AavgR, AavgG, AavgB uint32
    var BavgR, BavgG, BavgB uint32
    for x := sx; x < fx; x += 1 {
        for y := sy; y < fy; y += 1 {
            var ar, ag, ab, _ = a.At(x, y).RGBA()
            var br, bg, bb, _ = b.At(x, y).RGBA()
            AavgR += ar
            AavgG += ag
            AavgB += ab
            BavgR += br
            BavgG += bg
            BavgB += bb
        }
    }

    return int(math.Sqrt( math.Pow(float64(AavgR - BavgR), 2) +
                          math.Pow(float64(AavgG - BavgG), 2) +
                          math.Pow(float64(AavgB - BavgB), 2) ) )
}

func getTileWithMinimalDiff(part image.Image, tiles []image.Image) image.Image {
    var min_diff int = 1e10
    var min_idx int = 0

    for idx, tile := range tiles {
        var curr_diff = diff(part, tile)
        if min_diff > curr_diff {
            min_diff = curr_diff
            min_idx = idx
        }
    }

    return tiles[min_idx]
}


func Compose(main_image image.Image, tiles []image.Image) image.Image {
    var bounds image.Rectangle = main_image.Bounds()
    var sx, sy, fx, fy = bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y

    var res [][]image.Image = make([][]image.Image, int(math.Floor(float64(fx - sx) / TILE_WIDTH)))
    var resulting_image = image.NewRGBA(bounds)

    for x := 0; x < int(math.Floor(float64(fx - sx) / TILE_WIDTH)); x++ {
        res[x] = make([]image.Image, int(math.Floor(float64(fy - sy) / TILE_HEIGHT)))
        for y := 0; y < int(math.Floor(float64(fy - sy) / TILE_HEIGHT)); y++ {
            var rect = image.Rect(x * TILE_WIDTH, y * TILE_HEIGHT, (x + 1) * TILE_WIDTH, (y + 1) * TILE_HEIGHT)

            var subimg = main_image.(interface {
                SubImage(r image.Rectangle) image.Image
            }).SubImage(rect)

            var tile = getTileWithMinimalDiff(subimg, tiles)
            draw.Draw(resulting_image, rect, tile, image.Point{0, 0}, draw.Src)
        }
    }

    return resulting_image
}