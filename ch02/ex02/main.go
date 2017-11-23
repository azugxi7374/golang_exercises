package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)


type Celsius float64
type Fahrenheit float64
type Meter float64
type Feet float64
type Pound float64
type KGram float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c)}
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f)}
func (m Meter) String() string { return fmt.Sprintf("%gm", m)}
func (ft Feet) String() string { return fmt.Sprintf("%gft", ft)}
func (lb Pound) String() string { return fmt.Sprintf("%glb", lb)}
func (kg KGram) String() string { return fmt.Sprintf("%gkg", kg)}

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func MToFt(m Meter) Feet { return Feet(m*3.28084) }
func FtToM(f Feet) Meter { return Meter(f*0.3048) }
func LbToKg(lb Pound) KGram { return KGram(lb*0.453592) }
func KgToLb(kg KGram) Pound { return Pound(kg*2.20462) }


func main(){
    var v float64
    var err error
    if len(os.Args) > 1 {
        v, err = strconv.ParseFloat(os.Args[1], 64)
    } else {
        input := bufio.NewScanner(os.Stdin)
        if input.Scan() {
            v, err = strconv.ParseFloat(input.Text(), 64)
        }
    }
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    cf, fc, mft, ftm, lbkg, kglb := run(v)
    fmt.Printf(
        "%s = %s, %s = %s\n" +
        "%s = %s, %s = %s\n" +
        "%s = %s, %s = %s\n",
        Celsius(v), cf, Fahrenheit(v), fc,
        Meter(v), mft, Feet(v), ftm,
        Pound(v), lbkg, KGram(v), kglb,
    )
}

func run(v float64) (Fahrenheit, Celsius, Feet, Meter, KGram, Pound) {
    c:= Celsius(v)
    f:= Fahrenheit(v)
    m:= Meter(v)
    ft:= Feet(v)
    lb:= Pound(v)
    kg:= KGram(v)

    return CToF(c), FToC(f),
    MToFt(m), FtToM(ft),
    LbToKg(lb), KgToLb(kg)
}
