# shippy-user-service/Makefile

build:
	protoc -I. --go_out=plugins=micro:. \
		proto/user/user.proto
	# GOOS=linux ARCH=amd64 go build
	docker build -t shippy-user-service .

run:
	docker run --net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_USER=postgres \
		-e DB_PASSWORD=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		shippy-user-service
