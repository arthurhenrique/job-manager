package cmd

import (
	"hasty-challenge-manager/api"

	"github.com/spf13/cobra"
)

var (
	apiCommand = &cobra.Command{
		Use:   "api",
		Short: "Initialize the hasty job manager",
		Long:  "Initialize the hasty job manager",
		RunE:  apiExecute,
	}
)

func init() {
	RootCmd.AddCommand(apiCommand)
}

func apiExecute(cmd *cobra.Command, args []string) error {
	return api.Setup()
}
