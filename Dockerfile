FROM golang:1.17.1

WORKDIR /app
COPY go.mod .
COPY go.mod .
RUN go mod download

COPY . .

CMD [ "go", "run", "main.go" ]

