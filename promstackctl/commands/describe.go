package commands

import "github.com/spf13/cobra"

var describeRootCmd = &cobra.Command{
	Use:   "describe",
	Short: "describe root command",
	Long:  `describe given resource from the API endpoints.`,
}

func init() {
	describeRootCmd.AddCommand(serversDescribeCmd)
}
