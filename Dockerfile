FROM golang:1.19 AS builder

WORKDIR $GOPATH/src/kubedump

COPY . ./

RUN go get -u

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kubedump .


FROM scratch

COPY --from=builder /go/src/kubedump/kubedump ./

ENTRYPOINT ["./kubedump"]