package main

import (
	"fmt"
	"math"
    "os"
    "strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)


func main() {
    funcs := []func(float64, float64) float64 {
        fSin,
        f1,
        f2,
        f3,
    }
    idx, _ := strconv.Atoi(os.Args[1])
    fmt.Print(createSVG(funcs[idx], true))
}

func createSVG(f func(float64, float64) float64, skipNan bool) string {
	svg := ""
	svg += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for jj := 0; jj < cells; jj++ {
            polygon, clr := createPolygon(i, jj, f)
			if !skipNan || isValidPolygon(polygon) {
                svg += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s;'/>\n",
					polygon[0], polygon[1], polygon[2], polygon[3], polygon[4], polygon[5], polygon[6], polygon[7], clr)
			}
		}
	}
	svg += fmt.Sprintln("</svg>")
	return svg
}

// NaN, Inf, -Inf
func isValidPolygon(points []float64) bool {
    ok := true
    for i:=0; i< len(points); i++ {
        ok = ok && !math.IsNaN(points[i]) && !math.IsInf(points[i], 0)
    }
    return ok
}

func createPolygon(i, jj int, f func(float64, float64) float64) ([]float64, string) {
    ax, ay := corner(i+1, jj, f)
    bx, by := corner(i, jj, f)
    cx, cy := corner(i, jj+1, f)
    dx, dy := corner(i+1, jj+1, f)

    gz := f(xyrange * ((float64(i)+0.5)/cells - 0.5), xyrange * ((float64(jj)+0.5)/cells - 0.5))
    clr := calcColor(gz)
    return []float64{ax, ay, bx, by, cx, cy, dx, dy}, clr
}

func corner(i, jj int, f func(float64, float64) float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(jj)/cells - 0.5)

	z := f(x, y)

    // fmt.Println("corner", i, jj, z, x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}


func calcColor(z float64) string {
    maxz := height / zscale / 2.0
    minz := -maxz
    nz := (z - minz) / (maxz - minz)
    clr := fmt.Sprintf("rgb(%d,%d,%d)", int(0xff*nz), 0, int(0xff*(1-nz)))
    // fmt.Println(nz, z, maxz, minz, clr)
    return clr
}


func fSin(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

func fTan(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Tan(r) / r
}

func f1(x, y float64) float64 {
    h := 0.25
    ln := 1.5

    return - math.Abs(math.Cos(x / ln)) * math.Abs(math.Cos(y / ln)) * h
}

func f2(x, y float64) float64 {
    h := 0.25
    ln := 1.5

    return (math.Cos(x / ln) * math.Cos(y / ln)) * h
}

func f3(x, y float64) float64 {
    a := 0.0025
    b := 0.005
    return a * x * x - b * y * y
}
