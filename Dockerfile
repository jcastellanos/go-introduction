# Builder image
FROM golang:1.18.5-alpine3.16 as builder
RUN echo "--- Start build builder image ---"
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN echo "Executing tests"
RUN CGO_ENABLED=0 GOOS=linux go test ./...
RUN echo "Building binary"
RUN CGO_ENABLED=0 GOOS=linux go build -a -o go-rest-api .
RUN echo "--- End build builder image ---"
# Runtime image
RUN echo "--- Start build running image ---"
FROM alpine:3.16.2
COPY --from=builder /build/go-rest-api .
EXPOSE 8080

ENTRYPOINT [ "./go-rest-api" ]
RUN echo "--- End build running image ---"