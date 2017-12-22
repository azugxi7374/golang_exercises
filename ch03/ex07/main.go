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
            img.Set(px, py, newton(z))
        }
    }
    png.Encode(os.Stdout, img)
}

const EPS = 1e-9

func newton(z complex128) color.RGBA {
    const iterations = 200

    for n := uint8(0); n<iterations; n++ {
        z -= (z - 1.0/(z*z*z)) / 4.0

        v := uint8(255 * (1.0 - float64(n) / float64(iterations)))
        if cmplx.Abs(1+0i - z) < EPS {
            return color.RGBA{v, 0, 0, 0xff}
        } else if cmplx.Abs(0+1i - z) < EPS {
            return color.RGBA{0, v, 0, 0xff}
        } else if cmplx.Abs(-1+0i - z) <EPS {
            return color.RGBA{0, 0, v, 0xff}
        } else if cmplx.Abs(0-1i - z) < EPS {
            return color.RGBA{v, v, v, 0xff}
        }
    }
    return color.RGBA{0,0,0,255}
}
