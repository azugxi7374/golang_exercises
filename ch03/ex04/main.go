package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
    "strconv"
    "net/url"
)

var funcs = []func(float64, float64) float64 {
    fSin,
    f1,
    f2,
    f3,
}

func main() {
	http.HandleFunc("/", handler)
    http.HandleFunc("/test", handlerPrintParam)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL)
    r.ParseForm()
    params := createParamsFromForm(r.Form)
    w.Header().Set("Content-Type", "image/svg+xml")
    surface(w, params)
}

func handlerPrintParam(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    params := createParamsFromForm(r.Form)
    fmt.Fprintf(w, "%v", params)
}

///////////////////////////////////////////////
type Params struct {
    funcno int
	width, height, cells int
	xyrange, xyscaleRate, zscaleRate, angleRate float64
}

func defaultParams() Params {
    return Params{
        0, 600, 320, 100, 30.0, 2, 0.4, 30,
    }
}

func createParamsFromForm(form url.Values) Params{
    p := defaultParams()
    for k, v := range form {
        switch k {
        case "funcno":
            p.funcno, _ = strconv.Atoi(v[0])
        case "width":
            p.width, _ = strconv.Atoi(v[0])
        case "height":
            p.height, _ = strconv.Atoi(v[0])
        case "cells":
            p.cells, _ = strconv.Atoi(v[0])
        case "xyrange":
            p.xyrange, _ = strconv.ParseFloat(v[0], 64)
        case "xyscaleRate":
            p.xyscaleRate, _ = strconv.ParseFloat(v[0], 64)
        case "zscaleRate":
            p.zscaleRate, _ = strconv.ParseFloat(v[0], 64)
        case "angleRate":
            p.angleRate, _ = strconv.ParseFloat(v[0], 64)
        }
    }
    return p
}


////////////////////////////////////
// surface
func surface(out io.Writer, p Params) {
    fmt.Fprint(out, createSVG(p, true))
}

func createSVG(p Params, skipNan bool) string {
    svg := ""
    svg += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
    "style='stroke: grey; fill: white; stroke-width: 0.7' "+
    "width='%d' height='%d'>", p.width, p.height)

    for i := 0; i < p.cells; i++ {
        for jj := 0; jj < p.cells; jj++ {
            polygon, clr := createPolygon(i, jj, p)
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

func createPolygon(i, jj int, p Params) ([]float64, string) {
    ax, ay := corner(i+1, jj, p)
    bx, by := corner(i, jj, p)
    cx, cy := corner(i, jj+1, p)
    dx, dy := corner(i+1, jj+1, p)

    gz := funcs[p.funcno](p.xyrange * ((float64(i)+0.5)/float64(p.cells) - 0.5), p.xyrange * ((float64(jj)+0.5)/float64(p.cells) - 0.5))
    clr := calcColor(gz, p)
    return []float64{ax, ay, bx, by, cx, cy, dx, dy}, clr
}

func corner(i, jj int, p Params) (float64, float64) {
    xyscale := float64(p.width) / p.xyscaleRate / p.xyrange
    zscale := float64(p.height) * p.zscaleRate

    x := p.xyrange * (float64(i)/float64(p.cells) - 0.5)
    y := p.xyrange * (float64(jj)/float64(p.cells) - 0.5)

    z := funcs[p.funcno](x, y)

    sx := float64(p.width)/2 + (x-y)*math.Cos(2 * math.Pi * (p.angleRate / 360))*xyscale
    sy := float64(p.height)/2 + (x+y)*math.Sin(2 * math.Pi * (p.angleRate / 360))*xyscale - z*zscale
    return sx, sy
}


func calcColor(z float64, p Params) string {
    maxz := 1.0/ (2.0 * p.zscaleRate)
    minz := -maxz
    nz := (z - minz) / (maxz - minz)
    clr := fmt.Sprintf("rgb(%d,%d,%d)", int(0xff*nz), 0, int(0xff*(1-nz)))
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
