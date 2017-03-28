package commands

import (
	"os"
	"strconv"

	consul "github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/api/prometheus"
)

type config struct {
	Consul     consulConfig     `json:"consul"`
	Prometheus prometheusConfig `json:"prometheus"`
	Verbose    bool             `json:"verbose"`
}

type consulConfig struct {
	Address    string `json:"address"`
	client     *consul.Client
	Datacenter string `json:"datacenter"`
	Schema     string `json:"schema"`
	Port       string `json:"port"`
}

type prometheusConfig struct {
	Address string `json:"address"`
	client  prometheus.Client
	Schema  string `json:"schema"`
	Port    string `json:"port"`
}

var commandCfg config

//TODO:
// look at Kelsey's confd config.go. Should be a way to use the following
// scenario to line up appropriate configs:
// ENV -> FILE (/etc/serv/config.json:$PWD/.serv.json) -> ARG
// seems to be flag.Changed() can be used

func initConfig() {
	PromStackCmd.PersistentFlags().StringVar(&commandCfg.Consul.Address, "consul.address", "localhost", "Address to Consul API. [env:PROMSTACK_CONSUL_ADDRESS]")
	PromStackCmd.PersistentFlags().StringVar(&commandCfg.Consul.Datacenter, "consul.datacenter", "promstack", "Datacenter to reference in Consul API. [env:PROMSTACK_CONSUL_DATACENTER]")
	PromStackCmd.PersistentFlags().StringVar(&commandCfg.Consul.Port, "consul.port", "8500", "Port to Consul API. [env:PROMSTACK_CONSUL_PORT]")
	PromStackCmd.PersistentFlags().StringVar(&commandCfg.Consul.Schema, "consul.schema", "http", "Schema to access Consul API on. [env:PROMSTACK_CONSUL_SCHEMA]")

	PromStackCmd.PersistentFlags().StringVar(&commandCfg.Prometheus.Address, "prometheus.address", "localhost", "Address to Prometheus. [env:PROMSTACK_PROMETHEUS_ADDRESS]")
	PromStackCmd.PersistentFlags().StringVar(&commandCfg.Prometheus.Port, "prometheus.port", "9090", "Port to Prometheus API. [env:PROMSTACK_PROMETHEUS_PORT]")
	PromStackCmd.PersistentFlags().StringVar(&commandCfg.Prometheus.Schema, "promtheus.schema", "http", "Schema to access Prometheus API on. [env:PROMSTACK_PROMETHEUS_SCHEMA]")

	PromStackCmd.PersistentFlags().BoolVar(&commandCfg.Verbose, "verbose", false, "Enable Verbose output of application. [env:PROMSTACK_VERBOSE]")

	checkEnvironment()
}

func checkEnvironment() {
	//-- Consul ENV Variables for Configuration
	if consulAddress := os.Getenv("PROMSTACK_CONSUL_ADDRESS"); len(consulAddress) > 0 {
		commandCfg.Consul.Address = consulAddress
	}
	if consulDatacenter := os.Getenv("PROMSTACK_CONSUL_DATACENTER"); len(consulDatacenter) > 0 {
		commandCfg.Consul.Datacenter = consulDatacenter
	}
	if consulPort := os.Getenv("PROMSTACK_CONSUL_PORT"); len(consulPort) > 0 {
		commandCfg.Consul.Port = consulPort
	}
	if consulSchema := os.Getenv("PROMSTACK_CONSUL_SCHEMA"); len(consulSchema) > 0 {
		commandCfg.Consul.Schema = consulSchema
	}
	//-- Prometheus ENV Variables for Configuration
	if prometheusAddress := os.Getenv("PROMSTACK_PROMETHEUS_ADDRESS"); len(prometheusAddress) > 0 {
		commandCfg.Prometheus.Address = prometheusAddress
	}
	if prometheusSchema := os.Getenv("PROMSTACK_PROMETHEUS_SCHEMA"); len(prometheusSchema) > 0 {
		commandCfg.Prometheus.Schema = prometheusSchema
	}
	if prometheusPort := os.Getenv("PROMSTACK_PROMETHEUS_PORT"); len(prometheusPort) > 0 {
		commandCfg.Prometheus.Port = prometheusPort
	}

	if verbose := os.Getenv("PROMSTACK_VERBOSE"); len(verbose) > 0 {
		var err error
		if commandCfg.Verbose, err = strconv.ParseBool(verbose); err != nil {
			commandCfg.Verbose = false
		}
	}
}
