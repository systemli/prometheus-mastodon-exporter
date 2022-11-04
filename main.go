package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr        = flag.String("web.listen-address", ":13120", "Address on which to expose metrics and web interface.")
	mastodonURL = flag.String("mastodon-url", "", "Url from the Mastodon Instance (e.g.: https://mastodon.social)")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if *mastodonURL == "" {
		log.Fatal("mastodon url is empty")
	}
	_, err := url.Parse(*mastodonURL)
	if err != nil {
		log.Fatal("unable to parse mastodon url")
	}

	prometheus.MustRegister(NewCollector())
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
