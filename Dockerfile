FROM node:16 as node

WORKDIR /app
COPY web .
RUN yarn install
RUN yarn build

FROM golang:1.20 as builder

WORKDIR /app
COPY . /app
RUN mkdir /app/bin

COPY --from=node /app/build/ /app/web/build/

RUN go build -o bin/psc api/cmd/psc/main.go

CMD ["/app/bin/psc"]

