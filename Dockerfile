# Build in a stock Go builder container
FROM golang:1.17-alpine as builder

RUN apk add --no-cache gcc musl-dev linux-headers git

ADD . /case
RUN cd /case && go mod download && go build

FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /case/paymentSys /usr/local/bin/

EXPOSE 5050
ENTRYPOINT ["paymentSys"]