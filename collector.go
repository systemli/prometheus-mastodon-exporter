package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type Instance struct {
	Stats struct {
		UserCount   int `json:"user_count"`
		StatusCount int `json:"status_count"`
		DomainCount int `json:"domain_count"`
	} `json:"stats"`
}

type Activities []Activity

type Activity struct {
	Statuses      string `json:"statuses"`
	Logins        string `json:"logins"`
	Registrations string `json:"registrations"`
}

type Collector struct {
	Users              *prometheus.Desc
	Statuses           *prometheus.Desc
	Domains            *prometheus.Desc
	WeeklyStatuses     *prometheus.Desc
	WeeklyLogins       *prometheus.Desc
	WeeklyRegistration *prometheus.Desc
}

func NewCollector() *Collector {
	labels := []string{"host"}
	return &Collector{
		Users:              prometheus.NewDesc("mastodon_users", "Total number of all users", labels, nil),
		Statuses:           prometheus.NewDesc("mastodon_statuses", "Total number of all statuses", labels, nil),
		Domains:            prometheus.NewDesc("mastodon_domains", "Total number of known domains", labels, nil),
		WeeklyStatuses:     prometheus.NewDesc("mastodon_weekly_statuses", "Total number of weekly published statuses", labels, nil),
		WeeklyLogins:       prometheus.NewDesc("mastodon_weekly_logins", "Total number of weekly logins", labels, nil),
		WeeklyRegistration: prometheus.NewDesc("mastodon_weekly_registrations", "Total number of weekly registration", labels, nil),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Users
	ch <- c.Statuses
	ch <- c.Domains
	ch <- c.WeeklyStatuses
	ch <- c.WeeklyLogins
	ch <- c.WeeklyRegistration
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	mastodon, _ := url.Parse(*mastodonURL)
	labels := []string{mastodon.Host}

	instance, err := fetchInstance()
	if err == nil {
		ch <- prometheus.MustNewConstMetric(c.Users, prometheus.GaugeValue, float64(instance.Stats.UserCount), labels...)
		ch <- prometheus.MustNewConstMetric(c.Statuses, prometheus.GaugeValue, float64(instance.Stats.StatusCount), labels...)
		ch <- prometheus.MustNewConstMetric(c.Domains, prometheus.GaugeValue, float64(instance.Stats.DomainCount), labels...)
	} else {
		log.Println(err)
	}
	activities, err := fetchActivity()
	if err == nil {
		statuses, err := strconv.Atoi(activities[0].Statuses)
		if err == nil {
			ch <- prometheus.MustNewConstMetric(c.WeeklyStatuses, prometheus.GaugeValue, float64(statuses), labels...)
		}
		logins, err := strconv.Atoi(activities[0].Logins)
		if err == nil {
			ch <- prometheus.MustNewConstMetric(c.WeeklyLogins, prometheus.GaugeValue, float64(logins), labels...)
		}
		registrations, err := strconv.Atoi(activities[0].Registrations)
		if err == nil {
			ch <- prometheus.MustNewConstMetric(c.WeeklyRegistration, prometheus.GaugeValue, float64(registrations), labels...)
		}
	} else {
		log.Println(err)
	}
}

func fetchInstance() (Instance, error) {
	var instance Instance
	res, err := http.Get(*mastodonURL + "/api/v1/instance")
	if err != nil {
		return instance, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&instance)
	return instance, err
}

func fetchActivity() (Activities, error) {
	var activities Activities
	res, err := http.Get(*mastodonURL + "/api/v1/instance/activity")
	if err != nil {
		return activities, err
	}
	defer res.Body.Close()
	
	err = json.NewDecoder(res.Body).Decode(&activities)
	return activities, err
}
