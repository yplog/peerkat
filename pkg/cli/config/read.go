package config

import (
	"fmt"
	"log"

	"github.com/yplog/peerkat/pkg/config"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Print config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		read, err := config.Read()
		if err != nil {
			log.Fatal(err)
		}

		if read != nil {
			fmt.Printf("Config file content: %+v\n", *read)
			return
		}

		fmt.Println("Config file not found")
	},
}

func init() {
	Cmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
