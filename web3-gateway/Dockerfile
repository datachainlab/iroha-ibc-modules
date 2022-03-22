FROM golang:1.17.8-alpine3.14 as builder

RUN apk update
RUN apk add --no-cache gcc libc-dev

WORKDIR /root

COPY ./ ./

WORKDIR /root/web3-gateway
RUN go build -v -mod=readonly -o ./build/web3-gateway ./

FROM alpine:3.14

WORKDIR /root

COPY --from=builder /root/web3-gateway/build/web3-gateway /usr/local/bin/web3-gateway

RUN chmod +x /usr/local/bin/web3-gateway

EXPOSE 8545

ENTRYPOINT ["web3-gateway"]

