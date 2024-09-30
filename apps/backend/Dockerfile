FROM golang:1.23.1-alpine AS builder

RUN apk add --no-cache upx

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o server -a -ldflags="-s -w" -installsuffix cgo

RUN upx --ultra-brute -qq server && upx --t server

FROM scratch

COPY --from=builder /app/server /server

ENTRYPOINT [ "/server" ]