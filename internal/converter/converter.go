package converter

import (
	"bytes"
	"image"
	_ "image/jpeg" // Register JPEG decoder
	_ "image/png"  // Register PNG decoder
	"io"
	"os"

	"github.com/disintegration/imaging"
	"github.com/thezmc/jp2b/internal/config"
	_ "golang.org/x/image/webp" // Register WebP decoder
)

// Converter handles the image to braille conversion
type Converter struct {
	config *config.Config
}

// NewConverter creates a new Converter with the given configuration
func NewConverter(cfg *config.Config) *Converter {
	return &Converter{
		config: cfg,
	}
}

// Convert processes the input image and converts it to braille ASCII art
func (c *Converter) Convert() (string, error) {
	// Load the image
	img, err := c.loadImage()
	if err != nil {
		return "", err
	}

	// Convert to grayscale
	grayImg := imaging.Grayscale(img)

	// Calculate dimensions
	srcWidth := grayImg.Bounds().Dx()
	srcHeight := grayImg.Bounds().Dy()

	// Calculate output dimensions
	outWidth, outHeight := c.calculateOutputDimensions(srcWidth, srcHeight)

	// Resize image to match output dimensions, considering each braille cell represents 2x4 pixels
	resizedImg := imaging.Resize(grayImg, outWidth*2, outHeight*4, imaging.Lanczos)

	// Convert to braille
	result, err := c.convertToBraille(resizedImg, outWidth, outHeight)
	if err != nil {
		return "", err
	}

	// Save to file if output path is specified
	if c.config.OutputPath != "" {
		err = os.WriteFile(c.config.OutputPath, []byte(result), 0o644)
		if err != nil {
			return "", err
		}
	}

	return result, nil
}

// loadImage loads the image from the input path
func (c *Converter) loadImage() (image.Image, error) {
	file, err := os.Open(c.config.InputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// calculateOutputDimensions calculates the dimensions of the output based on config
func (c *Converter) calculateOutputDimensions(imgWidth, imgHeight int) (int, int) {
	// If both width and height are specified, use them
	if c.config.Width > 0 && c.config.Height > 0 {
		return c.config.Width, c.config.Height
	}

	// Calculate aspect ratio
	aspectRatio := float64(imgWidth) / float64(imgHeight) * 2.0 // Multiply by 2 because braille cells are 2x4

	// If only width is specified, calculate height
	if c.config.Width > 0 {
		height := max(int(float64(c.config.Width)/aspectRatio), 1)
		return c.config.Width, height
	}

	// If only height is specified, calculate width
	if c.config.Height > 0 {
		width := max(int(float64(c.config.Height)*aspectRatio), 1)
		return width, c.config.Height
	}

	// If neither is specified, use a default width of 80 characters
	defaultWidth := 80
	defaultHeight := max(int(float64(defaultWidth)/aspectRatio), 1)
	return defaultWidth, defaultHeight
}

// convertToBraille converts the grayscale image to braille ASCII art
func (c *Converter) convertToBraille(img image.Image, width, height int) (string, error) {
	var buf bytes.Buffer

	for y := range height {
		for x := range width {
			// Create a 2x4 grid for each braille character
			var brailleGrid [2][4]bool

			// Populate the grid based on pixel values
			for dy := range 4 {
				for dx := range 2 {
					posX := x*2 + dx
					posY := y*4 + dy

					// Get pixel grayscale value
					gray := getGrayValue(img, posX, posY)

					// Apply threshold
					isDot := gray > c.config.Threshold

					// Apply inversion if needed
					if c.config.Invert {
						isDot = !isDot
					}

					brailleGrid[dx][dy] = isDot
				}
			}

			// Create and write the braille character
			char := CreateBrailleChar(brailleGrid)
			_, err := buf.WriteRune(char)
			if err != nil {
				return "", err
			}
		}
		// Add a newline after each row
		_, err := buf.WriteRune('\n')
		if err != nil {
			return "", err
		}
	}

	return buf.String(), nil
}

// getGrayValue returns the grayscale value (0.0-1.0) of a pixel
func getGrayValue(img image.Image, x, y int) float64 {
	bounds := img.Bounds()
	if x < bounds.Min.X || x >= bounds.Max.X || y < bounds.Min.Y || y >= bounds.Max.Y {
		return 0.0 // Out of bounds
	}

	// Convert to grayscale value between 0.0 and 1.0
	r, g, b, _ := img.At(x, y).RGBA()
	gray := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 65535.0
	return gray
}

// ConvertToFile converts the image and writes the result to the specified writer
func (c *Converter) ConvertToFile(writer io.Writer) error {
	result, err := c.Convert()
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(result))
	return err
}
