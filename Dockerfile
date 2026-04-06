# Builder stage
FROM debian:stable-slim AS builder
ENV PATH="/root/.local/bin:${PATH}"

WORKDIR /app

RUN apt-get update && apt-get install -y curl git
RUN curl https://mise.run | sh

COPY .mise.toml .

RUN mise settings experimental=true && \
  mise trust && \
  mise i

COPY . .

RUN mise r install && \
  mise r test && \
  mise r build


FROM debian:stable-slim
ENV DATA="/app/data/"
ENV CORS="*"
ENV LOG_LEVEL="INFO"

COPY --from=builder /app/build/bin /app/
WORKDIR /app

CMD ["sh", "-c", "/app/bin -log-level=${LOG_LEVEL} -data=${DATA} -cors=${CORS}"]

