FROM golang:1.20.2-alpine3.17

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o main ./main.go

FROM alpine:3.17
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=0 /app/main ./
CMD ["/app/main"]

