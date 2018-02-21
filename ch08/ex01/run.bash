TZ=US/Eastern go run clock.go -port 8000&
TZ=Asia/Tokyo go run clock.go -port 8001&
TZ=Europe/London go run clock.go -port 8002&

go run clockwall.go NewYork=localhost:8000 Tokyo=localhost:8001 London=localhost:8002
