FROM golang:1.24.0-alpine as development

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app cmd/webserver/main.go

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=development /go/bin/app /go/bin/app

EXPOSE 8080

ENTRYPOINT ["/go/bin/app"]