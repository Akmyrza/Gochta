package main

import (
	"io"
	"log"
	"os"

	"github.com/cheggaaa/pb/v3"
)

func main() {

	var offset int64 = 10
	var limit int64 = 1024 * 1024 * 100
	
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Can't open input file!")
	}

	outputFile, err := os.Create("output.txt")
	if err!= nil {
		log.Fatal("Can't open output file!")
	}
	
	reader := io.NewSectionReader(inputFile, offset, limit)
	writer := outputFile

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(reader)

	io.Copy(writer, barReader)

	bar.Finish()

	defer outputFile.Close()
	defer inputFile.Close()
}