# Two-stage build:
#    first  FROM prepares a binary file in full environment ~780MB
#    second FROM takes only binary file ~10MB

FROM golang:1.12 AS builder

RUN go version

COPY . "/go/src/github.com/wlun/movie-api"
WORKDIR "/go/src/github.com/wlun/movie-api"

RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

RUN CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o /main

CMD ["/main"]

EXPOSE 5000

########
# second stage to obtain a very small image
#FROM scratch
#
#COPY --from=builder /go/src/github.com/wlun/movie-api/.env .
#COPY --from=builder /main .
#
#EXPOSE 5000
#
#CMD ["/main"]
