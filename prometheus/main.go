package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	systems      prometheus.Gauge
	systems_dupl prometheus.Gauge
	upgrades     *prometheus.CounterVec
	duration     *prometheus.HistogramVec
}

func newMetrics(reg prometheus.Registerer) *metrics {
	p := promauto.With(reg) // Factory that creates metrics with the custom registry provided
	namespace := "myapp"

	m := &metrics{
		systems: p.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "running_systems",
			Help:      "Number of running systems",
		}),
		systems_dupl: p.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "running_systems",
			Help:      "Number of running systems",
		}),
		upgrades: p.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "system_upgrade_total",
			Help:      "Total number of device upgrades",
		}, []string{"name"}),
		duration: p.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "request_duration_seconds",
			Help:      "Duration of requests",
			// Buckets define the upper bounds of the histogram buckets.
			Buckets: []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		}, []string{"method", "path"}),
	}

	fmt.Println(m.systems)
	fmt.Println(m.systems_dupl)

	return m
}

type systems struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

var sys []systems

func init() {
	sys = []systems{
		{Name: "system1", Version: "1.0"},
		{Name: "system2", Version: "1.1"},
		{Name: "system3", Version: "1.2"},
	}
}

func sleep(maxMS int) {
	time.Sleep(time.Duration(rand.Intn(maxMS)) * time.Millisecond)
}

func main() {
	// We are using a custom registry to avoid the default registry that is used by the promhttp.Handler and contains the Go runtime metrics.
	reg := prometheus.NewRegistry()
	// If you want to add the Go runtime metrics to the custom registry, you can do it like this:
	// reg.MustRegister(collectors.NewGoCollector())
	m := newMetrics(reg)

	m.systems.Set(float64(len(sys)))

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	http.Handle("/metrics", promHandler)

	http.HandleFunc("GET /systems", getSystems(m))
	http.HandleFunc("POST /systems", createSystem(m))
	http.HandleFunc("PATCH /systems/{name}", upgradeSystem(m))

	fmt.Println("Server started at :7766")
	http.ListenAndServe(":7766", nil)
}

func getSystems(m *metrics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// This type of metric is used in a middleware, I'm doing it here for simplicity
		now := time.Now()
		defer func() {
			m.duration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(now).Seconds())
		}()

		b, err := json.Marshal(sys)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sleep(300)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func createSystem(m *metrics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var s systems
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sys = append(sys, s)
		m.systems.Set(float64(len(sys))) // Gauges perform better with set operations than with inc/dec

		w.WriteHeader(http.StatusCreated)
	}
}

func upgradeSystem(m *metrics) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		systemName := r.PathValue("name")

		var s systems
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, v := range sys {
			if v.Name == systemName {
				sys[i].Version = s.Version
				break
			}
		}

		sleep(1000)

		m.upgrades.With(prometheus.Labels{"name": systemName}).Inc()
		w.WriteHeader(http.StatusOK)
	}
}
