package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
		if flagVerbose && flagQuiet {
			return fmt.Errorf("cannot use --verbose and --quiet together")
		}
		if len(args) == 0 {
			return fmt.Errorf("no files provided")
		}
		f := strings.ToLower(flagFormat)
		if f != "text" && f != "json" && f != "csv" {
			return fmt.Errorf("invalid format: %s", flagFormat)
		}
		flagFormat = f

		return nil
	},
}

var (
	flagLines   int
	flagFormat  string
	flagVerbose bool
	flagQuiet   bool
)

type Stats struct{ Lines, Words, Chars int }

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

func countFile(path string, maxLines int) (Stats, error) {
	f, err := os.Open(path)
	if err != nil {
		return Stats{}, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	stats := Stats{}
	for scanner.Scan() {
		line := scanner.Text()
		stats.Lines++
		stats.Words += len(strings.Fields(line))
		stats.Chars += len(line)
		if maxLines > 0 && stats.Lines >= maxLines {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return Stats{}, err
	}
	return stats, nil

}
