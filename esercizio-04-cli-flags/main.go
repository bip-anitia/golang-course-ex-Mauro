package main

import (
	"os"

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

var (
	flagLines   int
	flagFormat  string
	flagVerbose bool
	flagQuiet   bool
)

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().IntVar(&flagLines, "lines", 0, "number of lines to process")
	countCmd.Flags().StringVar(&flagFormat, "format", "text", "output format")
	countCmd.Flags().BoolVar(&flagVerbose, "verbose", false, "verbose output")
	countCmd.Flags().BoolVar(&flagQuiet, "quiet", false, "quiet output")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
