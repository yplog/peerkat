package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

// healthCmd represents the health command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Health is a network tool to test connectivity",
	Long:  `It sends a request to the peerkat server and returns the health result.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("health called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	Cmd.AddCommand(healthCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
