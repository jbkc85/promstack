package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

type exporter struct {
	Name string   `json:"name"`
	Port int      `json:"port"`
	Tags []string `json:"tags"`
}

var exporterGetCmd = &cobra.Command{
	Use:   "exporters",
	Short: "list available metadata for exporters",
	Long:  `list available metadata for exporters from the Consul KV store.`,
	Run:   getExporters,
}

func getExporters(cmd *cobra.Command, args []string) {
	pairs := getKVPath("promstack/exporters")

	table := uitable.New()
	table.Separator = "\t|\t"
	table.Wrap = true

	table.AddRow("EXPORTER", "PORT", "TAGS")
	table.AddRow("")

	for _, pair := range pairs {
		kvExporter := exporter{}
		if err := json.Unmarshal(pair.Value, &kvExporter); err != nil {
			log.Printf("[ERROR] Unable to unmarshal %v into structure, message: %s.", pair.Value, err)
		} else {
			table.AddRow(strings.Replace(pair.Key, "promstack/exporters/", "", -1), kvExporter.Port, kvExporter.Tags)
		}
	}
	fmt.Println(table)
}
