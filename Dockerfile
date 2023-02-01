FROM golang:1.20-alpine3.16 AS builder

WORKDIR /hasty-challenge-manager

COPY . .

RUN go mod download && go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build --ldflags='-w -s -extldflags "-static"' -v -a -o /go/bin/hasty-challenge-manager .

FROM alpine:3.17

COPY --from=builder /go/bin/hasty-challenge-manager /go/bin/hasty-challenge-manager

EXPOSE 9000

ENTRYPOINT ["/go/bin/hasty-challenge-manager"]
