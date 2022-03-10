FROM golang:alpine as builder

WORKDIR /app
COPY ./main.go /app/main.go

RUN go build -o a.out main.go

FROM alpine as runner

WORKDIR /app

COPY --from=builder /app .

CMD [ "./a.out" ]
# CMD ["go","run","main.go"]