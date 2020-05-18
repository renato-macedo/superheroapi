FROM golang:1.14.3-alpine3.11

RUN mkdir /api

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

ADD . /api


RUN go build

CMD [ "./superheroapi" ]