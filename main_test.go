package main

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"testing"
)

func TestReadImage(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create a grayscale image
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	white := color.Gray{255}
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, white)
		}
	}

	// Define the output path
	outputPath := filepath.Join(tempDir, "output.jpg")

	// Write the image to the output path
	file, err := os.Create(outputPath)
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	if err := jpeg.Encode(file, img, nil); err != nil {
		t.Fatalf("Failed to write image to output file: %v", err)
	}
	file.Close()

	// Call ReadImage with the output path
	readImg, err := imageprocessing.ReadImage(outputPath)
	if err != nil {
		t.Fatalf("ReadImage returned an error: %v", err)
	}

	// Check that the read image is equal to the original image
	if !imagesAreEqual(img, readImg) {
		t.Errorf("Read image is not equal to the original image")
	} else {
		t.Log("SUCCESS")
	}
}
func TestWriteImage(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create a grayscale image
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	white := color.Gray{255}
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, white)
		}
	}

	// Define the output path
	outputPath := filepath.Join(tempDir, "output.jpg")

	// Call WriteImage with the image and the output path
	err := imageprocessing.WriteImage(outputPath, img)
	if err != nil {
		t.Fatalf("WriteImage returned an error: %v", err)
	}

	// Check that the output file was created
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Output file was not created")
	}

	// Check that the output file is a valid JPEG image
	file, err := os.Open(outputPath)
	if err != nil {
		t.Fatalf("Failed to open output file: %v", err)
	}
	defer file.Close()
	if _, err := jpeg.Decode(file); err != nil {
		t.Errorf("Output file is not a valid JPEG image: %v", err)
	} else {
		t.Log("SUCCESS")
	}
}

func TestGrayscale(t *testing.T) {
	// Create a colored image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	red := color.RGBA{255, 0, 0, 255}
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, red)
		}
	}

	// Call Grayscale with the image
	grayImg := imageprocessing.Grayscale(img)

	// Check that the image was converted to grayscale
	bounds := grayImg.Bounds()
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			r, g, b, _ := grayImg.At(x, y).RGBA()
			if r != g || g != b {
				t.Errorf("Pixel at (%d, %d) was not converted to grayscale", x, y)
			}
		}
	}
}

func TestResize(t *testing.T) {
	// Create a colored image
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	red := color.RGBA{255, 0, 0, 255}
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			img.Set(x, y, red)
		}
	}

	// Define the new width and height
	newWidth, newHeight := 500, 500

	// Call Resize with the image
	resizedImg := imageprocessing.Resize(img)

	// Check that the image was resized correctly
	if resizedImg.Bounds().Dx() != newWidth || resizedImg.Bounds().Dy() != newHeight {
		t.Errorf("Image was not resized correctly, got: %dx%d, want: %dx%d", resizedImg.Bounds().Dx(), resizedImg.Bounds().Dy(), newWidth, newHeight)
	} else {
		t.Log("SUCCESS")
	}
}

// Helper function to test equality
func imagesAreEqual(img1, img2 image.Image) bool {
	if img1.Bounds() != img2.Bounds() {
		return false
	}
	for y := 0; y < img1.Bounds().Dy(); y++ {
		for x := 0; x < img1.Bounds().Dx(); x++ {
			if img1.At(x, y) != img2.At(x, y) {
				return false
			}
		}
	}
	return true
}

func BenchmarkPipeline(b *testing.B) {
	imagePaths := []string{"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	for i := 0; i < b.N; i++ {
		channel1, _ := loadImage(imagePaths)
		channel2 := resize(channel1)
		channel3 := convertToGrayscale(channel2)
		writeResults := saveImage(channel3)

		for success := range writeResults {
			if !success {
				b.Fatal("Failed!")
			}
		}
	}
}
