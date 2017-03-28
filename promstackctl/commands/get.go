package commands

import "github.com/spf13/cobra"

var getRootCmd = &cobra.Command{
	Use:   "get",
	Short: "get root command",
	Long:  `get given resource from the API endpoints.`,
}

func init() {
	getRootCmd.AddCommand(exporterGetCmd)
	getRootCmd.AddCommand(serversGetCmd)
}
