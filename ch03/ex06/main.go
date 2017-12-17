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
        y0 := float64(py) / height * (ymax - ymin) + ymin
        y1 := (float64(py)+0.5) / height * (ymax - ymin) + ymin

        for px := 0; px < width; px++ {
            x0 := float64(px) / width * (xmax - xmin) + xmin
            x1 := (float64(px)+0.5) / width * (xmax - xmin) + xmin

            comps := []complex128{ complex(x0, y0), complex(x1, y0), complex(x0, y1), complex(x1, y1)}
            r, g, b := 0, 0, 0
            for i:=0; i<4; i++ {
                m := mandelbrot(comps[i])
                r += int(m.R)
                g += int(m.G)
                b += int(m.B)
            }
            img.Set(px, py, color.RGBA{uint8(r/4), uint8(g/4), uint8(b/4), 255})
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
