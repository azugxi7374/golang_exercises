package main

import (
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "time"
)

var palette = []color.Color{
    color.RGBA{0x00, 0x00, 0x00, 0xff},
    color.RGBA{0xff, 0x00, 0x00, 0xff}, // red
    color.RGBA{0xff, 0x7f, 0x00, 0xff}, // orange
    color.RGBA{0xff, 0xff, 0x00, 0xff}, // yellow
    color.RGBA{0x00, 0xff, 0x00, 0xff}, // green
    color.RGBA{0x00, 0xff, 0xff, 0xff}, // cyan
    color.RGBA{0x00, 0x00, 0xff, 0xff}, // blue
    color.RGBA{0xff, 0x00, 0xff, 0xff}, // purple
}

const clrBackIdx = 0
var clrLineIdx = []uint8{1, 2, 3, 4, 5, 6, 7}


func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles = 5
        res = 0.001
        size = 100
        nframes = 64
        delay = 8
    )
    freq := rand.Float64() * 3.0
    anim := gif.GIF{ LoopCount: nframes }
    phase := 0.0
    for i := 0; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for x := -size; x <= size; x++ {
            for y := -size; y <= size; y++ {
                img.SetColorIndex(size + x, size + y, clrBackIdx)
            }
        }
        tMax := cycles*2*math.Pi
        for t := 0.0; t < tMax; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            clri := clrLineIdx[int(t / tMax * float64(len(clrLineIdx)))]
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), clri)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
