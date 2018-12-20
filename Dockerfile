# Builder image, where we build, lint and test the binary
FROM golang:latest AS builder
ARG GitlabAPIURL
ARG GitlabAPIToken
ENV GO111MODULE=on

WORKDIR /go/src/rettungsring

COPY . .

RUN go get -u golang.org/x/lint/golint
RUN golint -set_exit_status ./...

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -o rettungsring

# Final image
FROM alpine:latest
COPY --from=builder /go/src/rettungsring/rettungsring /usr/local/bin/rettungsring
RUN apk add --no-cache ca-certificates

WORKDIR /data

ENTRYPOINT [ "/usr/local/bin/rettungsring" ]
CMD [ "" ]
