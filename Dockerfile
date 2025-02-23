FROM golang:1.23 as builder
WORKDIR /build
COPY . ./

RUN go mod download
RUN CGO_ENABLED=0 go build -a -o ./.bin/app ./cmd/main.go

FROM alpine:latest
COPY --from=builder /build/.bin/app /bin/app

ENTRYPOINT ["/bin/app"]