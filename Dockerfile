# Copies the binary produced by `mise run build` (build/bin). Tests and the
# build itself run in CI via mise; Docker only packages the deliverable.
FROM gcr.io/distroless/static-debian12:nonroot
ENV PORT=8000 \
    DATA=/app/data \
    CORS=* \
    LOG_LEVEL=INFO
COPY build/bin /app/bin
EXPOSE 8000
ENTRYPOINT ["/app/bin"]
