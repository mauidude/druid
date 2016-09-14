FROM golang:1.6.0-alpine

ENV GOPATH /opt/go
WORKDIR /opt/go

ADD . /opt/go

RUN go build -v event_pump

ENV KAFKA_BROKER_LIST=kafka:9092

CMD ["./event_pump"]
