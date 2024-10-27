FROM golang:1.18-bullseye
COPY . /app
WORKDIR /app
RUN go build -o app main.go

FROM debian:bullseye-slim
WORKDIR /root
COPY --from=0 /app/app ./
ENTRYPOINT ["./app"]