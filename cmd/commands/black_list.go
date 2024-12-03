package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var blackListCmd = &cobra.Command{
	Use:   "blacklist",
	Short: "Manage blacklist networks",
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println(cmd.UsageString())
	},
}

func init() {
	rootCmd.AddCommand(blackListCmd)
}
