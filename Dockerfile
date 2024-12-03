FROM golang:1.22-alpine3.20 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /anti-brute-force-app /build/internal/app/app.go

FROM alpine:3.20

COPY --from=builder /build/configs/config.yaml /configs/config.yaml

COPY --from=builder anti-brute-force-app /bin/anti-brute-force-app

EXPOSE 9012

CMD [ "/bin/anti-brute-force-app" ]
