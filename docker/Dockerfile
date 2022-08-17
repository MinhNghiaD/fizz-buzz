FROM golang:1.18 as builder
ENV GO111MODULE=on
WORKDIR /opt/fizzbuzz
COPY . .
RUN make build

FROM scratch
WORKDIR /opt/fizzbuzz
COPY --from=builder /opt/fizzbuzz/bin/server .
ENTRYPOINT [ "./server" ]