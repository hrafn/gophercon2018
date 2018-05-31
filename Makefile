PROJECT?=github.com/hrafn/gophercon2018
APP?=gophercon
PORT?=8000

clean:
	rm -f ./bin/${APP}

build: clean
	#CG0_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH}
	go build \
		-o ./bin/${APP} ${PROJECT}./cmd/gophercon/

run: build
	PORT=${PORT} ./bin/${APP}

test:
	go test -race ./..
