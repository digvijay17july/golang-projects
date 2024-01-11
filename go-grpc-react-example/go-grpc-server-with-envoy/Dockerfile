FROM golang:1.21.4-alpine

WORKDIR /app

COPY go.sum ./


COPY . ./

RUN go build -o  /go-grpc-server-with-envoy

EXPOSE 8080

CMD [ "/go-grpc-server-with-envoy" ]