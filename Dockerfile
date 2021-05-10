FROM golang:1.16.3 as builder

WORKDIR /work
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN cd cmd/toy-isuumo &&\
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /work/toy-isuumo

FROM alpine:3.13
COPY --from=builder /work/toy-isuumo /usr/local/bin
ENTRYPOINT ["toy-isuumo"]