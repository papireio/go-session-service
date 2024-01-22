FROM golang:1.21

WORKDIR /usr/src
COPY . .
RUN make vendor \
    && make build \
    && mkdir -p /usr/app \
    && cp ./target/goservice /usr/app

WORKDIR /usr/app
RUN rm -rf /usr/src

CMD ["./goservice"]
