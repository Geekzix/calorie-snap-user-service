FROM golang:1.27.3-alpine AS builder

WORKDIR /src

COPY go.mod ./
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /api ./cmd/server/

FROM scratch

COPY --from=builder /api /api

USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/api"]

