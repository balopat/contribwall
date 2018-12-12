FROM golang:1.10.1-alpine3.7

WORKDIR /go/src/github.com/balopat/contribwall/backend/svc-contributors
ADD . .
RUN go build -o out/svc-contributors svc/contributors/contributors.go

CMD ["./out/svc-contributors"]
