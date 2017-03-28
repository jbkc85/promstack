package commands

import (
	"log"

	consul "github.com/hashicorp/consul/api"
)

//-- initializing functions
func initializeConsul() {
	consulConfig := &consul.Config{
		Address:    commandCfg.Consul.connectionString(),
		Datacenter: commandCfg.Consul.Datacenter,
	}

	consulClient, err := consul.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("[ERROR] Unable to connect to Consul, message: %s", err)
	} else {
		commandCfg.Consul.client = consulClient
	}
}

//-- getNodes
func getNodes() []*consul.Node {
	catalog := commandCfg.Consul.client.Catalog()
	nodes, _, err := catalog.Nodes(&consul.QueryOptions{})
	if err != nil {
		log.Printf("[ERROR] Unable to retrieve nodes from Consul Catalog, message: %s", err)
	}
	return nodes
}

//-- getNode
func getNode(nodeID string) *consul.CatalogNode {
	catalog := commandCfg.Consul.client.Catalog()
	node, _, err := catalog.Node(nodeID, &consul.QueryOptions{})
	if err != nil {
		log.Printf("[ERROR] Unable to retrieve node from Consul Catalog, message: %s", err)
	}
	return node
}

//-- getKV
func getKV(path string) *consul.KVPair {
	kv := commandCfg.Consul.client.KV()
	value, meta, err := kv.Get(path, nil)
	if err != nil {
		log.Fatalf("[ERROR] Unable to retrieve KV list from promstack/exporters")
		return &consul.KVPair{}
	}
	if commandCfg.Verbose {
		log.Printf("[DEBUG] Meta: %v", meta)
	}
	return value
}

//-- getKVPath
func getKVPath(path string) consul.KVPairs {
	kv := commandCfg.Consul.client.KV()
	pairs, meta, err := kv.List(path, nil)
	if err != nil {
		log.Fatalf("[ERROR] Unable to retrieve KV list from promstack/exporters")
		return nil
	}

	if commandCfg.Verbose {
		log.Printf("[DEBUG] Meta: %v", meta)
	}
	return pairs
}

//-- struct/object functions
func (c *consulConfig) connectionString() string {
	return c.Address + ":" + c.Port
}

func (c *consulConfig) health() (string, error) {
	var health = "ok"
	catalog := c.client.Catalog()
	_, _, err := catalog.Nodes(&consul.QueryOptions{})

	if err != nil {
		log.Printf("[ERROR] Consul Health Check Failed, message: %s", err)
		health = "error"
	}
	return health, err
}
