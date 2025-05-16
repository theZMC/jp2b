# jp2b

jp2b (JPEG to Braille) is a command-line tool written in Go that converts images
(PNG/JPG/WEBP) to ASCII art using Unicode braille characters (U+2800 to U+28FF).

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
go build -o jp2b ./cmd/jp2b

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
1. The image is resized to match the desired output dimensions.
1. For each 2x4 block of pixels, a braille character is generated:
   - Each pixel in the block corresponds to a dot in the braille pattern.
   - If a pixel's brightness is above the threshold, the corresponding dot is
     activated.
1. The resulting braille characters are arranged in rows and columns to form the
   ASCII art.

## License

MIT License

## AI Assistance

This tool was created with the assistance of Anthropic's Claude 3.7 Sonnet. See
[the AI assistance](./AI-ASSISTANCE.md) document for more information.
