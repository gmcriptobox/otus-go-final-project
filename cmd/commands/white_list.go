package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var whiteListCmd = &cobra.Command{
	Use:   "whitelist",
	Short: "Manage whitelist networks",
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println(cmd.UsageString())
	},
}

func init() {
	rootCmd.AddCommand(whiteListCmd)
}
