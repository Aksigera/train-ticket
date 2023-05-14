# syntax=docker/dockerfile:1

FROM golang

WORKDIR /app

ADD ./main.go  /app/
ADD ./go.mod /app/

RUN go build
CMD [ "./app" ]

EXPOSE 12862
