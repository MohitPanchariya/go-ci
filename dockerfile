# Build stage
FROM golang:1.22-bullseye AS build

WORKDIR /app

COPY ./go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o go-ci .

# Final stage
FROM alpine:edge

WORKDIR /app

COPY --from=build /app/go-ci .

CMD ["./go-ci"]
