FROM  golang:1.12-alpine3.10 As builder

ENV GO111MODULE=on

ENV GOPROXY=https://goproxy.cn

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /shortener-core

ADD . /shortener-core

RUN go mod tidy

RUN go build main.go

FROM alpine:3.10

WORKDIR /shortener-core

COPY --from=builder /shortener-core .

COPY config ./config

CMD ["/shortener-core/main"]

