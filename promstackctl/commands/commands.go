package commands

import "github.com/spf13/cobra"

// PromStackCmd ...
var PromStackCmd = &cobra.Command{
	Use:           "promstackctl",
	Short:         "promstackctl is a command line script to interact with the PromStack Monitoring Suite.",
	Long:          `PromStack is a collection of software solutions to simplify ones introduction into monitoring, logging and alerting by providing a pre-configured stack of default software around Prometheus.`,
	SilenceErrors: true,
	SilenceUsage:  true,
}

// where should this go...

var monitorRootCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitor",
	Long:  `monitor`,
}

func init() {
	monitorRootCmd.AddCommand(serversMonitorCmd)
	PromStackCmd.AddCommand(healthRootCmd)
	PromStackCmd.AddCommand(getRootCmd)
	PromStackCmd.AddCommand(describeRootCmd)
	PromStackCmd.AddCommand(monitorRootCmd)

	initConfig()

	// initialize Consul
	initializeConsul()
	initializePrometheus()
}
