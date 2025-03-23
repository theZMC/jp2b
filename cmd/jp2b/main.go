package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thezmc/jp2b/internal/config"
	"github.com/thezmc/jp2b/internal/converter"
)

var cfg = &config.Config{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jp2b",
	Short: "Convert images to ASCII art using braille characters",
	Long: `jp2b is a CLI tool that converts images (PNG/JPG/WEBP) to ASCII art
using Unicode braille characters (U+2800 to U+28FF). The resulting ASCII art
can be displayed in the terminal or saved to a file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate configuration
		if err := cfg.Validate(); err != nil {
			return err
		}

		// Create converter
		conv := converter.NewConverter(cfg)

		// Convert the image
		result, err := conv.Convert()
		if err != nil {
			return err
		}

		// If output path is not specified, print to stdout
		if cfg.OutputPath == "" {
			fmt.Print(result)
			return nil
		}

		// Otherwise, save to file
		err = os.WriteFile(cfg.OutputPath, []byte(result), 0o644)
		if err != nil {
			return err
		}

		fmt.Printf("ASCII art saved to %s\n", cfg.OutputPath)
		return nil
	},
}

func init() {
	// Define command line flags
	rootCmd.Flags().StringVarP(&cfg.InputPath, "input", "i", "", "Input image path (required)")
	rootCmd.Flags().StringVarP(&cfg.OutputPath, "output", "o", "", "Output text file path (defaults to stdout)")
	rootCmd.Flags().IntVarP(&cfg.Width, "width", "w", 0, "Output width in braille characters")
	// Use "H" instead of "h" for height to avoid conflict with help
	rootCmd.Flags().IntVarP(&cfg.Height, "height", "H", 0, "Output height in braille characters")
	rootCmd.Flags().Float64VarP(&cfg.Threshold, "threshold", "t", 0.5, "Brightness threshold (0.0-1.0)")
	rootCmd.Flags().BoolVarP(&cfg.Invert, "invert", "v", false, "Invert black and white")

	// Mark required flags
	rootCmd.MarkFlagRequired("input")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
