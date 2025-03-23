FROM golang:1.24 AS builder

WORKDIR $GOPATH/src/kubedump

COPY . ./

RUN go get -u

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kubedump .


FROM cgr.dev/chainguard/wolfi-base:latest

COPY --from=builder /go/src/kubedump/kubedump ./

ENTRYPOINT ["./kubedump"]