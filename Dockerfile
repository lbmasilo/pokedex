FROM golang:1.12-alpine
RUN apk add --no-cache git

WORKDIR /pokedex-api-v1

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/pokedex-api-v1 .

EXPOSE 8080


CMD ["./out/pokedex-api-v1"]