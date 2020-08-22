FROM golang:1.14-alpine

RUN apk add --no-cache git

WORKDIR $GOPATH/src/github.com/berto/kerbal

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o ./kerbal .
EXPOSE 3000:3000
CMD ["./kerbal"]