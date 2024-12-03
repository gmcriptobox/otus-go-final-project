package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "abf",
	Short: "Anti-brute-force cli tool",
	Run: func(cmd *cobra.Command, _ []string) {
		fmt.Println(cmd.UsageString())
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("ip", "", serviceAddr, "service address")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
