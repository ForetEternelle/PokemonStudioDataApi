FROM golang:1.23.2 AS builder

RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY . /src/
WORKDIR /src/
RUN task build

FROM debian:bookworm
ENV DATA="/app/data/"
ENV CORS="*"
ENV LOG_LEVEL="INFO"

COPY --from=builder /src/build/psapi /app/
WORKDIR /app

CMD ["sh", "-c", "/app/psapi -log-level=${LOG_LEVEL} -data=${DATA} -cors=${CORS}"]

