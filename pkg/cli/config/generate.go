package config

import (
	"fmt"
	"log"

	"github.com/yplog/peerkat/pkg/config"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate default config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		exist := config.Exists()
		if exist {
			var input string

			fmt.Print("Config file already exists, overwrite? (y/n) ")
			fmt.Scanln(&input)

			if input != "y" {
				return
			}
		}

		generate, err := config.Generate()
		if err != nil {
			log.Fatal(err)
			return
		}

		if generate != "" {
			fmt.Println("Config file generated:", generate)
			return
		}
	},
}

func init() {
	Cmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
