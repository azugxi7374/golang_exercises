fetched=`go run fetch.go https://golang.org`
want=`echo $fetched | go run main_original.go`
got=`echo $fetched | go run main.go`

if [ "$want" = "$got" ] ; then
    echo "ok"
else
    echo $want
    echo ""
    echo $got
    echo "failed"
fi
