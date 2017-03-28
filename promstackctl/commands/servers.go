package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gosuri/uitable"
	consul "github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
)

var serversGetCmd = &cobra.Command{
	Use:   "servers",
	Short: "get servers from Consul Catalog",
	Long:  `get a list of Servers registered to the Consul Catalog.`,
	Run:   getServers,
}

var serversDescribeCmd = &cobra.Command{
	Use:   "servers",
	Short: "describe a server from Consul Catalog",
	Long:  `describe a specific Server registered in the Consul Catalog.`,
	Run:   describeServer,
}

var serversMonitorCmd = &cobra.Command{
	Use:   "servers",
	Short: "monitor a given server",
	Long:  `monitor a given server from a predefined exporter in the Consul Catalog.`,
	Run:   monitorServer,
}

var (
	nodeNameFlag     string
	nodeAddressFlag  string
	exporterNameFlag string
)

func init() {
	serversMonitorCmd.PersistentFlags().StringVar(&nodeAddressFlag, "node.address", "", "IPv4 Address (or DNS if website) of node to monitor")
	serversMonitorCmd.PersistentFlags().StringVar(&nodeNameFlag, "node.name", "", "name of node to monitor")
	serversMonitorCmd.PersistentFlags().StringVar(&exporterNameFlag, "exporter.name", "", "name of exporter to monitor on node")
}

func getServers(cmd *cobra.Command, args []string) {
	catalog := commandCfg.Consul.client.Catalog()

	consulChan := make(chan []*consul.Node)
	go func() {
		nodes, _, err := catalog.Nodes(&consul.QueryOptions{})
		if err != nil {
			log.Printf("[ERROR] Unable to connect to Consul Catalog, message: %s", err)
		}
		consulChan <- nodes
	}()

	nodes := <-consulChan

	table := uitable.New()
	table.Separator = "\t|\t"
	table.Wrap = true
	table.AddRow("SERVER", "ADDRESS")
	table.AddRow("")
	for _, n := range nodes {
		table.AddRow(n.Node, n.Address)
	}
	fmt.Println(table)
}

func describeServer(cmd *cobra.Command, args []string) {
	if args[0] == "" {
		log.Printf("[ERROR] Unable to describe Server, please use the server in question as an argument to the script.")
	} else {
		node := getNode(args[0])

		table := uitable.New()
		table.Wrap = true
		table.AddRow("SERVER", node.Node.Node)
		table.AddRow("ADDRESS", node.Node.Address)
		table.AddRow("Services")
		for _, s := range node.Services {
			table.AddRow("    " + s.Service + " [" + node.Node.Address + ":" + strconv.Itoa(s.Port) + "]")
		}
		fmt.Println(table)
	}
}

func monitorServer(cmd *cobra.Command, args []string) {
	if nodeAddressFlag == "" {
		log.Fatalf("[ERROR] Please use the --node.address argument to provide a node address to monitor.")
		return
	}
	if nodeNameFlag == "" {
		log.Fatalf("[ERROR] Please use the --node.name argument to provide a node name to monitor.")
		return
	}
	if exporterNameFlag == "" {
		log.Fatalf("[ERROR] Please use the --exporter.name argument to provide an exporter to monitor.")
		return
	}

	exporterDetails := getKV("promstack/exporters/" + exporterNameFlag)

	if exporterDetails.Value == nil {
		log.Fatalf("[ERROR] Unable to retrieve Exporter from Consul Inventory.  Please try again.")
	} else {
		kvExporter := exporter{}
		if err := json.Unmarshal(exporterDetails.Value, &kvExporter); err != nil {
			log.Printf("[ERROR] Unable to unmarshal %v into structure, message: %s.", exporterDetails.Value, err)
		} else {
			newNode := &consul.CatalogRegistration{
				Node:    nodeNameFlag,
				Address: nodeAddressFlag,
				Service: &consul.AgentService{
					Service: exporterNameFlag,
					Port:    kvExporter.Port,
					Tags:    kvExporter.Tags,
				},
			}

			catalog := commandCfg.Consul.client.Catalog()

			meta, err := catalog.Register(newNode, &consul.WriteOptions{})

			if err != nil {
				log.Printf("[ERROR] Unable to Register %s to Catalog, message: %s", nodeNameFlag, err)
			}

			log.Printf("[DEBUG] Catalog Meta: %v", meta)

		}
	}

}
