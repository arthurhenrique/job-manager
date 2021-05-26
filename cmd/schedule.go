package cmd

import (
	"hasty-challenge-manager/schedule"

	"github.com/spf13/cobra"
)

var (
	scheduleCheckerCommand = &cobra.Command{
		Use:   "schedule-checker",
		Short: "Run the hasty schedule checker",
		Long:  "Run the hasty schedule checker",
		RunE:  scheduleCheckerExecute,
	}
)

func init() {
	RootCmd.AddCommand(scheduleCheckerCommand)
}

func scheduleCheckerExecute(cmd *cobra.Command, args []string) error {
	return schedule.Run()
}
