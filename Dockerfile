# Builder image
FROM golang:1.18.5-alpine3.16 as builder
RUN mkdir /build
ADD * /build/
WORKDIR /build
#RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -o go-rest-api .

# Runtime image
FROM alpine:3.16.2
COPY --from=builder /build/go-rest-api .
EXPOSE 8080

ENTRYPOINT [ "./go-rest-api" ]