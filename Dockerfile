FROM alpine:3.23.2 as builder

WORKDIR /go/src/github.com/systemli/prometheus-mastodon-exporter

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN apk add --no-cache --update ca-certificates

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY prometheus-mastodon-exporter /prometheus-mastodon-exporter

USER appuser:appuser

EXPOSE 13120

ENTRYPOINT ["/prometheus-mastodon-exporter"]
