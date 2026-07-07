FROM docker.arvancloud.ir/golang:1.24.5-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -o /out/demo-api \
    ./cmd/server

FROM docker.arvancloud.ir/alpine:3.22
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app

COPY --from=build --chown=app:app /out/demo-api /app/demo-api
USER app
EXPOSE 8080
ENV PORT=8080
ENV APP_VERSION=v1

ENTRYPOINT ["/app/demo-api"]
