/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tavvfiq/go-kafka-testbed/internal/app"
)

// producerCmd represents the producer command
var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "start application as a producer",
	Run: func(cmd *cobra.Command, args []string) {
		app.StartProducer()
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// producerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// producerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
