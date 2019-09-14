FROM golang:1.13

WORKDIR ./src

COPY . /opt
EXPOSE 8088

RUN go build /opt/src/server.go

CMD ["./server"]

