# AI Assistance Notice

This project was developed with assistance from Anthropic's Claude 3.7 Sonnet, an AI language model.

## Contributions from AI

The AI assisted with:

1. Initial code structure and organization
2. Implementation of the core image processing and braille conversion logic
3. Creating the command-line interface using Cobra
4. Handling proper image format support, including WebP
5. Troubleshooting flag conflicts and file output issues

## Replication Prompt

The following prompt was used to create the initial version of this tool. This can be useful if you want to:

- Understand how the project was conceived
- Create similar tools with AI assistance
- Extend or modify this tool with further AI assistance

````
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

## Sample Usage

```bash
# Basic usage (outputs to terminal)
$ brailleimg -i image.png

# Specify width, auto-calculate height
$ brailleimg -i image.jpg -w 80

# Specify both dimensions
$ brailleimg -i image.webp -w 100 -h 50 -o output.txt

# Adjust threshold and invert
$ brailleimg -i image.png -t 0.7 -v
```
````

## Best Practices for AI-Assisted Development

When using AI to help with software development:

1. **Review all code**: Understand what the AI has generated and make sure it meets your requirements and standards.

2. **Test thoroughly**: AI may generate code that compiles but has subtle bugs or edge cases.

3. **Add human creativity**: Enhance AI-generated code with your own ideas and improvements.

4. **Provide clear prompts**: The quality of AI output depends greatly on the clarity and specificity of your instructions.

5. **Acknowledge AI usage**: Be transparent about AI assistance in your projects, as done here.
