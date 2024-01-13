package config

import (
	"fmt"
	"log"

	"github.com/yplog/peerkat/pkg/config"

	"github.com/spf13/cobra"
)

// pathCmd represents the path command
var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print config path",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := config.Path()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Config file path:", path)
	},
}

func init() {
	Cmd.AddCommand(pathCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pathCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pathCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
