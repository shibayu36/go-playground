# バイナリを作成するコンテナ
FROM golang:1.21.3-bookworm as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# --------------------

# デプロイ用
FROM debian:bookworm-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# --------------------

FROM golang:1.21.3-bookworm as dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]
