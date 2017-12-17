package main

import (
    "image"
    "image/color"
    "image/png"
    "math/cmplx"
    "os"
)

func main() {
    const (
        xmin, ymin, xmax, ymax = -2, -2, +2, +2
        width, height = 1024, 1024
    )

    img := image.NewRGBA(image.Rect(0, 0, width, height))
    for py := 0; py < height; py++ {
        y := float64(py) / height * (ymax - ymin) + ymin
        for px := 0; px < width; px++ {
            x := float64(px) / width * (xmax - xmin) + xmin
            z := complex(x, y)
            img.Set(px, py, mandelbrot(z))
        }
    }
    png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.RGBA {
    const iterations = 200

    var v complex128
    for n := uint8(0); n<iterations; n++ {
        v = v*v + z
        if cmplx.Abs(v) > 2 {
            return toColor(float64(n) / float64(iterations))
        }
    }
    return color.RGBA{0,0,0,255}
}

func toColor(d float64) color.RGBA{
    colors:= [][]int{
        {0x7f, 0, 0x7f},
        {0, 0, 0xff},
        {0, 0xff, 0xff},
        {0, 0xff, 0},
        {0xff, 0xff, 0},
        {0xff, 0xa5, 0},
        {0xff, 0, 0},
    }
    d *= float64(len(colors) - 1)
    di := int(d)
    dd := d - float64(int(d))

    r := float64(colors[di][0]) * (1-dd) + float64(colors[di+1][0]) * dd
    g := float64(colors[di][1]) * (1-dd) + float64(colors[di+1][1]) * dd
    b := float64(colors[di][2]) * (1-dd) + float64(colors[di+1][2]) * dd
    return color.RGBA{uint8(r),uint8(g),uint8(b),0xff}
}
