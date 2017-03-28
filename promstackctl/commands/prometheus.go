package commands

import (
	"context"
	"log"
	"time"

	"github.com/prometheus/client_golang/api/prometheus"
)

//-- initializing functions
func initializePrometheus() {
	prometheusClient, err := prometheus.New(prometheus.Config{
		Address: commandCfg.Prometheus.connectionString(),
	})
	if err != nil {
		log.Fatalf("[ERROR] Unable to connect to Prometheus API, message: %s", err)
	} else {
		commandCfg.Prometheus.client = prometheusClient
	}
}

func (p *prometheusConfig) connectionString() string {
	return p.Schema + "://" + p.Address + ":" + p.Port
}

func (p *prometheusConfig) health() (string, error) {
	var health = "ok"

	queryAPI := prometheus.NewQueryAPI(p.client)
	_, err := queryAPI.Query(context.Background(), "up{job='prometheus'}", time.Now())

	if err != nil {
		log.Printf("[ERROR] Prometheus Health Check Failed, message: %s", err)
		health = "error"
	}
	return health, err
}
