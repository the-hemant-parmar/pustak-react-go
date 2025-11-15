# minimal Dockerfile
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN apk add --no-cache git
RUN go build -o /book-service ./cmd/book-service


FROM alpine:3.18
COPY --from=build /book-service /book-service
EXPOSE 8081
CMD ["/book-service"]