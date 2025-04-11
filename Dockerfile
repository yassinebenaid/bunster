FROM golang:1.23

WORKDIR /bunster

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /usr/local/bin/bunster ./cmd/bunster

RUN rm -rf /bunster /tmp/* /var/tmp/*

CMD ["bash"]
