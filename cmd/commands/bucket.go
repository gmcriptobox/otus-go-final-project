package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Manage buckets",
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println(cmd.UsageString())
	},
}

func init() {
	rootCmd.AddCommand(bucketCmd)
}
