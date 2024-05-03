package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"os"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) (<-chan Job, <-chan error) {
	out := make(chan Job)
	errc := make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errc)
		for _, p := range paths {
			// Check if the input file exists and is readable
			if _, err := os.Stat(p); os.IsNotExist(err) {
				errc <- fmt.Errorf("input file %s does not exist", p)
				return
			} else if err != nil {
				errc <- fmt.Errorf("error reading input file %s: %v", p, err)
				return
			}

			// Prepare the output path
			outPath := strings.Replace(p, "images/", "images/output/", 1)

			// Check if the output path is writable
			outFile, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				errc <- fmt.Errorf("error opening output file %s for writing: %v", outPath, err)
				return
			}
			outFile.Close()

			job := Job{InputPath: p, OutPath: outPath}

			// Load the image
			job.Image, err = imageprocessing.ReadImage(p)
			if err != nil {
				errc <- fmt.Errorf("error loading image from %s: %v", p, err)
				return
			}

			out <- job
		}
	}()
	return out, errc
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {

	imagePaths := []string{"images/image1.jpeg",
		"images/image2.jpeg",
		"images/image3.jpeg",
		"images/image4.jpeg",
	}

	channel1, _ := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
}
