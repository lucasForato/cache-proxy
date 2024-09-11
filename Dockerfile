FROM golang:1.23

WORKDIR /usr/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 3000

RUN go build -o app .

ENTRYPOINT ["./app"]
