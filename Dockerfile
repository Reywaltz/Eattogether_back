FROM golang:1.23.1 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin ./cmd/main.go


FROM builder
COPY --from=builder /app/bin /server
ENTRYPOINT [ "/server" ]