# shippy-user-service/Makefile

build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/seiji-thirdbridge/shippy-user-service \
		proto/user/user.proto
	GOOS=linux ARCH=amd64 go build
	# docker build -t user-service .

