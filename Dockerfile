FROM alpine:latest

ADD output/app-linux_amd64 /root/app
RUN chmod +x /root/app
RUN apk add ca-certificates

WORKDIR /root

ENTRYPOINT [ "./app" ]


ADD config /root/config