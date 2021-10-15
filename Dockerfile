# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /gRPCMicroserviceswithGo

COPY go.mod /gRPCMicroserviceswithGo
COPY go.sum /gRPCMicroserviceswithGo
RUN go mod download

COPY . /gRPCMicroserviceswithGo

RUN go build -o /docker-gs-ping

EXPOSE 50051

CMD [ "/docker-gs-ping" ]

