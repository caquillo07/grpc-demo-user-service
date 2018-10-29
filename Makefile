generate:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/caquillo07/grpc-demo-user-service \
    	proto/user/user.proto

build: generate docker-image

docker-image:
	docker build -t user-service .

run:
	docker run --net="host" \
		-p 50051 \
		-e DB_HOST=localhost:5433 \
		-e DB_PASSWORD=root \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		user-service
