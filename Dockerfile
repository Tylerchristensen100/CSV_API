FROM golang:1.24.2 AS build

WORKDIR /app


COPY . /app


RUN go mod download && go mod verify

RUN go build -ldflags="-w" -o server .


ENV ENV=prod

EXPOSE 3000
CMD ["./server"]
