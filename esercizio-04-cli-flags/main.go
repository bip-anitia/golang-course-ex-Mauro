package main

import (
	"flag"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filetools",
	Short: "A versatile file processing tool",
}

var countCmd = &cobra.Command{
	Use:   "count [files...]",
	Short: "Count lines, words, and characters",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: implement
		return nil
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}

func main() {
	// TODO: Definire e parsare i flags

	flag.Parse()

	fmt.Println("CLI Tool con Flags")
}
