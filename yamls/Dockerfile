# builder
FROM golang:1.19.7 as builder
WORKDIR /usr/src/app
COPY main.go /usr/src/app/main.go
RUN go build main.go
# run
FROM scratch
COPY --from=builder /usr/src/app/main /
CMD ["./main"]
