FROM golang:1.26.5-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /booking-api ./cmd/server

FROM alpine:3.22
COPY --from=build /booking-api /booking-api
EXPOSE 8080
ENTRYPOINT ["/booking-api"]
