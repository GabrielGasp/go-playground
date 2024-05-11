FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY main.go ./

RUN CGO_ENABLED=0 go build -o /my-app



FROM gcr.io/distroless/base-debian11

COPY --from=builder /my-app /my-app

ENTRYPOINT ["/my-app"]