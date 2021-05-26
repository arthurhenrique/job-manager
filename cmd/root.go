package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "hasty-challenge-manager",
	Short: "hasty-challenge-manager - Hasty Backend Challenge",
	Long:  "hasty-challenge-manager - Hasty Backend Challenge",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(-1)
	}
}
