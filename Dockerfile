FROM golang:1.20 as builder

WORKDIR /app
COPY . /app
RUN mkdir /app/bin

RUN go build -o bin/psc api/cmd/psc/main.go

CMD ["/app/bin/psc"]

