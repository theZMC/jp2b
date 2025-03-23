# jp2b

jp2b (JPEG to Braille) is a command-line tool written in Go that converts images (PNG/JPG/WEBP) to ASCII art using Unicode braille characters (U+2800 to U+28FF).

## Features

- Convert PNG, JPG, and WEBP images to braille ASCII art
- Customize output width and height
- Adjust brightness threshold
- Invert black and white if needed
- Output to terminal or save to a file

## Installation

```bash
# Clone the repository
git clone https://github.com/thezmc/jp2b.git
cd jp2b

# Build the binary
go build -o jp2b ./cmd/brailleimg

# Optional: Move to a directory in your PATH
mv jp2b /usr/local/bin/
```

## Usage

```bash
# Basic usage (outputs to terminal)
jp2b -i image.png

# Specify width, auto-calculate height
jp2b -i image.jpg -w 80

# Specify both dimensions
jp2b -i image.webp -w 100 -h 50 -o output.txt

# Adjust threshold and invert
jp2b -i image.png -t 0.7 -v
```

## Command Line Arguments

- `--input, -i`: Input image path (required)
- `--output, -o`: Output text file path (optional, defaults to stdout)
- `--width, -w`: Output width in braille characters (optional)
- `--height, -H`: Output height in braille characters (optional)
- `--threshold, -t`: Brightness threshold (0.0-1.0, default 0.5)
- `--invert, -v`: Invert black and white (flag, default false)

## How It Works

1. The input image is loaded and converted to grayscale.
2. The image is resized to match the desired output dimensions.
3. For each 2x4 block of pixels, a braille character is generated:
   - Each pixel in the block corresponds to a dot in the braille pattern.
   - If a pixel's brightness is above the threshold, the corresponding dot is activated.
4. The resulting braille characters are arranged in rows and columns to form the ASCII art.

## License

MIT License

## AI Assistance

This tool was created with the assistance of Anthropic's Claude 3.7 Sonnet. The prompt below can be used to recreate or expand upon this tool:

```
Create a Golang application that converts images (PNG/JPG/WEBP) to ASCII art using only braille characters (Unicode range U+2800 to U+28FF).

## Application Structure

Use a multi-package structure:
- `cmd/`: Contains the cobra CLI implementation
- `internal/converter/`: Image processing and braille conversion logic
- `internal/config/`: Configuration types and validation

## Core Features

1. **Image Processing**:
   - Read PNG/JPG/WEBP images using Go's standard imaging libraries
   - Convert images to grayscale
   - Apply binary thresholding (black/white) based on configurable threshold

2. **Braille Mapping**:
   - Use a 2x4 grid of pixels to map to each braille character
   - Map each pixel in the grid to a dot in the braille pattern
   - Generate appropriate Unicode braille character based on dot pattern

3. **Sizing and Aspect Ratio**:
   - Calculate output dimensions based on input parameters
   - Maintain aspect ratio when only width or height is specified
   - Allow custom width and height that may distort the image if needed

## Command Line Arguments

Implement using Cobra with these flags:
- `--input, -i`: Input image path (required)
- `--output, -o`: Output text file path (optional, defaults to stdout)
- `--width, -w`: Output width in braille characters (optional)
- `--height, -h`: Output height in braille characters (optional)
- `--threshold, -t`: Brightness threshold (0.0-1.0, default 0.5)
- `--invert, -v`: Invert black and white (flag, default false)
```
