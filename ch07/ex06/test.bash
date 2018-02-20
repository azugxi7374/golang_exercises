function test() {
    P_TEMP=$1
    echo "$P_TEMP -> " `go run main.go -temp $P_TEMP`
}

test "273.15K"
test "0K"
test "1000K"

