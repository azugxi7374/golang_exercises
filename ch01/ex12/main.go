package main

import (
    // "fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
    "strconv"
	// "os"
	// "time"
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
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL)
    p := defaultLParam()
    r.ParseForm()
    for k, v := range r.Form {
        switch k {
        case "cycles":
            p.cycles, _ = strconv.Atoi(v[0])
        case "size":
            p.size, _ = strconv.Atoi(v[0])
        case "nframes":
            p.nframes, _ = strconv.Atoi(v[0])
        case "delay":
            p.delay, _ = strconv.Atoi(v[0])
        case "res":
            p.res, _ = strconv.ParseFloat(v[0], 64)
        }
    }
    lissajous(w, p)
}

///////////////////////////////////////////////
// Lissajous
type LissajousParam struct {
    cycles  int
    size    int
    nframes int
    delay   int
    res     float64
}

func defaultLParam() LissajousParam {
    return LissajousParam{
        5, 100, 64, 8, 0.001,
    }
}

func lissajous(out io.Writer, p LissajousParam) {
    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: p.nframes}
    phase := 0.0
    for i := 0; i < p.nframes; i++ {
        rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
        img := image.NewPaletted(rect, palette)
        for x := -p.size; x <= p.size; x++ {
            for y := -p.size; y <= p.size; y++ {
                img.SetColorIndex(p.size+x, p.size+y, clrBackIdx)
            }
        }
        tMax := float64(p.cycles) * 2 * math.Pi
        for t := 0.0; t < tMax; t += p.res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            clri := clrLineIdx[int(t/tMax*float64(len(clrLineIdx)))]
            img.SetColorIndex(p.size+int(x*float64(p.size)+0.5), p.size+int(y*float64(p.size)+0.5), clri)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, p.delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
