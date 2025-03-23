package config

import (
	"errors"
	"os"
	"path/filepath"
)

// Config represents the application configuration
type Config struct {
	InputPath  string  // Path to the input image
	OutputPath string  // Path to the output text file (empty means stdout)
	Width      int     // Output width in braille characters
	Height     int     // Output height in braille characters
	Threshold  float64 // Brightness threshold (0.0-1.0)
	Invert     bool    // Whether to invert black and white
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	// Check if input file exists
	if c.InputPath == "" {
		return errors.New("input path is required")
	}

	info, err := os.Stat(c.InputPath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("input path must be a file, not a directory")
	}

	// Check file extension
	ext := filepath.Ext(c.InputPath)
	switch ext {
	case ".png", ".jpg", ".jpeg", ".webp":
		// Valid extensions
	default:
		return errors.New("input file must be PNG, JPG, or WEBP")
	}

	// Validate threshold range
	if c.Threshold < 0.0 || c.Threshold > 1.0 {
		return errors.New("threshold must be between 0.0 and 1.0")
	}

	// Validate width and height if specified
	if c.Width < 0 {
		return errors.New("width must be positive")
	}
	if c.Height < 0 {
		return errors.New("height must be positive")
	}

	return nil
}
