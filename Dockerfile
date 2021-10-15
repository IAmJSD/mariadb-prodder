FROM golang:1.17-alpine
WORKDIR /var/app
COPY . .
RUN go build -o app

FROM alpine
RUN apk add ca-certificates
WORKDIR /var/app
COPY --from=0 /var/app/app /var/app/app
ENTRYPOINT ./app
