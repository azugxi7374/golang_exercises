package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "testing"
)

func TestLissajousParam(t *testing.T){
    var tests = []struct{
        param string
        want LissajousParam
    }{
        {"", LissajousParam{5,100,64,8,0.001}},
        {"cycles=10&res=0.05&delay=5", LissajousParam{10,100,64,5,0.05}},
        {"size=10&nframes=32", LissajousParam{5,10,32,8,0.001}},
    }
    for _, test := range tests {
        resp, _ := http.Get("http://localhost:8000/test?" + test.param)
        b, _ := ioutil.ReadAll(resp.Body)
        resp.Body.Close()

        got := fmt.Sprintf("%s", b)
        wantstr := fmt.Sprintf("%v", test.want)
        if wantstr != got {
            t.Errorf("Failed. input:%s, want:%s, got:%s\n", test.param, wantstr, got)
        }
    }
}
