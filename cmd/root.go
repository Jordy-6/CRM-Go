package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "crm-go",
	Short: "A simple CRM application written in Go",
	Long:  `crm-go is a command-line application that helps manage customer relationships efficiently.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define flags and configuration settings.
	// eg: rootCmd.Flags().StringP("config", "c", "", "Path to the config file")
}
