FROM golang:1.9.4 as builder
WORKDIR /go/src/app
COPY . .
RUN go-wrapper download
RUN go-wrapper install
RUN go build -o seppo main.go

FROM golang:1.9.4
COPY --from=builder /go/src/app/seppo /usr/bin
RUN chmod +x /usr/bin/seppo
CMD ["seppo"]