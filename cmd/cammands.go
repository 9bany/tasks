package cmd

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "start server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task --help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
