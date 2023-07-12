BUILD=out

build:
	go build -o ${BUILD} main/main.go 

run: build
	./${BUILD}
