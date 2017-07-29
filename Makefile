default: build

build:
	GOOS=0 go build -o herd.darwin main.go

