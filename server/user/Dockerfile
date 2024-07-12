# Build
FROM golang:1.22 as builder

WORKDIR /app

COPY . .
RUN go mod download

ENV CGO_ENABLED=1
ENV GOOS=linux
RUN go build -o user-service ./cmd/api/main.go

# Test
FROM builder as test
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian12 AS release

WORKDIR /
VOLUME ["/database"]

COPY --from=builder ./app/user-service .

ENV USER_DB_PATH="/database/"
ENV TOKEN_DB_PATH="/database/"
ENV PORT=5000
EXPOSE ${PORT}

USER nonroot:nonroot

ENTRYPOINT ["./content-service"]
