FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/America/Recife /etc/localtime && \
    echo "America/Recife" >  /etc/timezone && \
    apk del tzdata

WORKDIR /marsover

COPY . .

RUN go mod download && go mod verify

RUN apk add --no-cache git make

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --generalInfo ./internal/http/server.go --output ./docs

RUN GOOS=linux GOARCH=amd64 go build --ldflags='-w -s -extldflags "-static"' -v -a -o /go/bin/marsover .

FROM alpine:3.17

RUN adduser -D user-manager
USER user-manager

COPY --from=builder /go/bin/marsover /usr/bin/marsover
COPY --from=builder /docs/swagger/swagger.json /docs/swagger/swagger.json

EXPOSE 9000

ENTRYPOINT ["/usr/bin/marsover"]