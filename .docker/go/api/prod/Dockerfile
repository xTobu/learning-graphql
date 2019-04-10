FROM golang:1.11.2 as builder
WORKDIR /root/go
COPY . .
RUN go build -o ApiServer


FROM debian:9-slim as stage
WORKDIR /root/
COPY --from=builder ./root/go/ApiServer ./
COPY --from=builder ./root/go/config ./config/

EXPOSE 3000

CMD ["/root/ApiServer"]