FROM golang:1.11.0 as builder

WORKDIR /go/src/github.com/caquillo07/grpc-demo-user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
WORKDIR /app
COPY --from=builder /go/src/github.com/caquillo07/grpc-demo-user-service/user-service .

CMD ["./user-service"]