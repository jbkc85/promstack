package commands

import (
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

var healthRootCmd = &cobra.Command{
	Use:   "health",
	Short: "Check Health of Endpoints for PromStack",
	Long:  `Check all endpoints in the PromStack Suite for connectivity.`,
	Run:   healthCheck,
}

func init() {
}

func healthCheck(cmd *cobra.Command, args []string) {
	table := uitable.New()
	table.Wrap = true

	table.AddRow("ENDPOINT", "|", "STATUS", "|", "MESSAGE")
	consulHealth, consulErr := commandCfg.Consul.health()
	table.AddRow(commandCfg.Consul.connectionString(), "|", consulHealth, "|", consulErr)
	prometheusHealth, prometheusErr := commandCfg.Prometheus.health()
	table.AddRow(commandCfg.Prometheus.connectionString(), "|", prometheusHealth, "|", prometheusErr)

	fmt.Println(table)
}
