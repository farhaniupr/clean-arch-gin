# Builder image
FROM golang:1.18-alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o . .

#Final Image
FROM alpine:3.13
RUN apk add bash
RUN apk add curl
# RUN apk add build-base
# RUN apk add inotify-tools
RUN apk --no-cache add ca-certificates tzdata \
    && cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime \
    && mkdir -p /var/www/html/golang
WORKDIR /var/www/html/golang
COPY --from=builder /app/clean-gin .
COPY startservice.sh .
COPY .env .
EXPOSE 8000
CMD ["./clean-gin", "app:serve"]
# RUN ./clean-gin app:serve/
# RUN ["chmod", "+x", "./startservice.sh"]
# CMD ["/bin/bash", "-c", "./startservice.sh"]