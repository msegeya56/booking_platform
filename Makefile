.PHONY: build
build:
	mkdir -p bin
	go build -o bin/booking_platform cmd/web/routes.go

.PHONY: run
run:
	go run cmd/web/routes.go
