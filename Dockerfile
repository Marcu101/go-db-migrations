FROM golang:1.23

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build

CMD ["./go-db-migrations"]