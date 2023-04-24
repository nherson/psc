FROM node:16 as node

WORKDIR /app
COPY web .
RUN yarn build

FROM golang:1.20 as builder

WORKDIR /app
COPY . /app
RUN mkdir /app/bin

RUN go build -o bin/psc api/cmd/psc/main.go

COPY --from=node /app/build/ /app/web/build/


CMD ["/app/bin/psc"]

