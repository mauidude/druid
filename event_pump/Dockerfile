FROM golang:1.6.0-alpine

ENV GOPATH /opt/go
WORKDIR /opt/go

ADD . /opt/go

RUN go build -v event_pump

ENTRYPOINT ["./entrypoint.sh"]
CMD ["./event_pump"]
