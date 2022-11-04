# prometheus-mastodon-exporter

[![Integration](https://github.com/systemli/prometheus-mastodon-exporter/actions/workflows/integration.yaml/badge.svg)](https://github.com/systemli/prometheus-mastodon-exporter/actions/workflows/integration.yaml)
[![Quality](https://github.com/systemli/prometheus-mastodon-exporter/actions/workflows/quality.yaml/badge.svg)](https://github.com/systemli/prometheus-mastodon-exporter/actions/workflows/quality.yaml)
[![Release](https://github.com/systemli/prometheus-mastodon-exporter/actions/workflows/release.yaml/badge.svg)](https://github.com/systemli/prometheus-mastodon-exporter/actions/workflows/release.yaml)

Prometheus Exporter for Mastodon written in Go.

## Usage

```shell
go install github.com/systemli/prometheus-mastodon-exporter@latest
$GOPATH/bin/prometheus-mastodon-exporter -mastodon-url=https://mastodon.social
```

### Commandline options

```text
  -mastodon-url string
        Url from the Mastodon Instance (e.g.: https://mastodon.social)
  -web.listen-address string
        Address on which to expose metrics and web interface. (default ":13120")
```

## Metrics

```text
# HELP mastodon_domains Total number of known domains
# TYPE mastodon_domains gauge
mastodon_domains{host="mastodon.social"} 27289
# HELP mastodon_statuses Total number of all statuses
# TYPE mastodon_statuses gauge
mastodon_statuses{host="mastodon.social"} 4.0143351e+07
# HELP mastodon_users Total number of all users
# TYPE mastodon_users gauge
mastodon_users{host="mastodon.social"} 817561
# HELP mastodon_weekly_logins Total number of weekly logins
# TYPE mastodon_weekly_logins gauge
mastodon_weekly_logins{host="mastodon.social"} 34851
# HELP mastodon_weekly_registrations Total number of weekly registration
# TYPE mastodon_weekly_registrations gauge
mastodon_weekly_registrations{host="mastodon.social"} 281
# HELP mastodon_weekly_statuses Total number of weekly published statuses
# TYPE mastodon_weekly_statuses gauge
mastodon_weekly_statuses{host="mastodon.social"} 80742
```

### Docker

```shell
docker run -p 13120:13120 systemli/prometheus-mastodon-exporter:latest -mastodon-url=https://mastodon.social
```

## License

GPLv3
