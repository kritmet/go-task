FROM golang:1.18-alpine3.17 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update && apk upgrade 

RUN mkdir /api
WORKDIR /api
ADD . /api

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.2
RUN swag init

RUN go test ./...
RUN go build -o /app/api .

FROM alpine:3.17

COPY --from=builder /app/api /api
COPY --from=builder /api/docs /docs
ADD /configs /configs

ENV TZ=Asia/Bangkok

EXPOSE 8000

ENTRYPOINT ["/api"]
