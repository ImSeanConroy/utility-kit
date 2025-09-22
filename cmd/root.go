package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "utility-kit",
	Short: "Utility kit is a versatile CLI utility tool",
	Long:  `Utility kit is a lightweight CLI "toolbox" that helps you with everyday tasks.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generatePasswordCmd)
	rootCmd.AddCommand(generateMealCmd)
}
