FROM golang:1.21

WORKDIR /usr/src
COPY . .
RUN make vendor \
    && make build \
    && mkdir -p /usr/app \
    && cp ./target/go-session /usr/app

WORKDIR /usr/app
RUN rm -rf /usr/src

CMD ["./go-session"]
