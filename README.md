# goroutinespipeline-projectrefuge
This repository is a fork of the [go_21_goroutines_pipeline](https://github.com/code-heim/go_21_goroutines_pipeline) repository. The original repository contains a Go program that processes images using a pipeline of goroutines. The program reads an image file, converts it to grayscale, inverts the colors, and then saves the modified image to a new file. The program uses a pipeline of goroutines to process the image data concurrently.

The goal of this project is to add error checking, unit tests, and benchmarking to the original program. The modifications will improve the reliability of the program, ensure that it functions correctly, and provide performance metrics for the pipeline.

## Prerequisites

- Clone the [GitHub repository](https://github.com/code-heim/go_21_goroutines_pipeline) for image processing
- Build and run the program in its original form (done but undocumented)

## Major Modifications
- Add [error checking](#error-checking) for image file input and output
- [Replace the four input image files](#image-replacement) with files of your choosing
- Add [unit tests](#unit-tests) to the code repository
- Add [benchmark methods](#benchmarking) for capturing pipeline throughput times
- [Build, test, and run](#documented-build-test-and-run) the pipeline program.
- Prepare a [complete README.md file](#goroutinespipeline-projectrefuge) documenting your work.

## Error Checking

The original program does not check for errors when opening the input image file or creating the output image file. This can lead to unexpected behavior if the files are not found or cannot be opened. To improve the reliability of the program, we will add error checking for these operations.

Full functionality can be found at [this commit](https://github.com/ryano0oceros/goroutinespipeline-projectrefuge/commit/2edb60b94b8a863bf71fc824da84f814d25ef09a#diff-2873f79a86c0d8b3335cd7731b0ecf7dd4301eb19a82ef7a1cba7589b5252261). Major modifications are found in lines 24-52 of `main.go`. The `ReadImage` and `WriteImage` functions in `image_processing.go` needed to be modified to return errors.

```go
func ReadImage(path string) (image.Image, error)
```

Additional housekeeping changes needed to be made in made to `main.go` to handle the error returns from `ReadImage` and `WriteImage` functions. (i.e. `go return out` modified to `go return out, errc`)


## Image Replacement

Full functionality can be found at [this commit](https://github.com/ryano0oceros/goroutinespipeline-projectrefuge/commit/221ed4c62ab3e875d66b434ceaf139fcdd3feb0f). Self-explanatory as the image files were simply replaced with new ones.

## Unit Tests

The original program does not have any unit tests. To ensure that the program functions correctly, I added unit tests for the `ReadImage`, `WriteImage`, `Grayscale`, and `Invert` functions in `image_processing.go`. These tests will verify that the functions produce the expected output for a given input.

I couldn't find a native way to test image equality in Go, so I wrote a simple helper function, `imagesAreEqual`, to compare the pixel values of two images. 

Full functionality can be found at [this commit](https://github.com/ryano0oceros/goroutinespipeline-projectrefuge/commit/5b3963434ae3ffe772d8f5b338e547914d28a914). 

### Usage

```bash
go test -v

=== RUN   TestReadImage
    main_test.go:49: SUCCESS
--- PASS: TestReadImage (0.01s)
=== RUN   TestWriteImage
    main_test.go:88: SUCCESS
--- PASS: TestWriteImage (0.01s)
=== RUN   TestGrayscale
--- PASS: TestGrayscale (0.00s)
=== RUN   TestResize
    main_test.go:137: SUCCESS
--- PASS: TestResize (0.00s)
PASS
ok      goroutines_pipeline     0.560s
```

## Benchmarking

The original program does not have any benchmarking. To measure the performance of the pipeline, I created the `BenchmarkPipeline` function in `main_test.go`. This function runs the pipeline with a given number of iterations and measures the time taken to process the images. The benchmarking results can help provide insights into the throughput of the pipeline and help identify bottlenecks.

Full functionality can be found at [this commit](https://github.com/ryano0oceros/goroutinespipeline-projectrefuge/commit/425710c38d8417542a23ff16f8b9247ddf1182e8)

### Usage

```bash
go test -bench=BenchmarkPipeline                                              

goos: darwin
goarch: arm64
pkg: goroutines_pipeline
BenchmarkPipeline-10                  21          59533448 ns/op
PASS
ok      goroutines_pipeline     1.838s
```

## Documented Build, Test, and Run

The `README.md` file includes information on the program's functionality, how to replace the input images, how to run the unit tests, and how to run the benchmarking tests.

Full functionality can be found at [this commit](https://github.com/ryano0oceros/goroutinespipeline-projectrefuge/commit/13bf5a77c7a7bf627f670d24e24d0b5b1524cac1)

### Usage

```bash
go run main.go

Success!
Success!
Success!
Success!
```
