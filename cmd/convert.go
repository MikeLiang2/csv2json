/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/MikeLiang2/csv2json/internal/converter"
	"github.com/spf13/cobra"
)

// The convert command to convert CSV files to JSON Lines format.
var convertCmd = &cobra.Command{
	Use:   "convert [input.csv] [output.jl]",
	Short: "Convert CSV file to JSON Lines format",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		outputPath := args[1]

		err := converter.ConvertCSVToJSONLines(inputPath, outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ Conversion successful!")
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
