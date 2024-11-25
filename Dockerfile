# syntax=docker/dockerfile:1

FROM golang:1.23.3 AS build-stage
  WORKDIR /app

  COPY go.mod go.sum ./
  RUN go mod download

  COPY . .

  RUN CGO_ENABLED=0 GOOS=linux go build -o /vitshop ./cmd/main.go

FROM scratch AS build-release-stage
  WORKDIR /

  COPY --from=build-stage /vitshop /vitshop

  EXPOSE 8080

  ENTRYPOINT ["/vitshop"]
