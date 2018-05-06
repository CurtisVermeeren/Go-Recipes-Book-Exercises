package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

// addToArchive writes a given file into a zip file
func addToArchive(filename string, zw *zip.Writer) error {
	// Open the given file to archive into a zip file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create addas a file to the zip file using the given name
	// returns a io.Writer to which file contents should be written
	wr, err := zw.Create(filename)
	if err != nil {
		return nil
	}
	// Write the file contents to the zip file
	if _, err := io.Copy(wr, file); err != nil {
		return err
	}

	return nil
}

// archiveFiles writes a group of files into a zip file
func archiveFiles(files []string, archive string) error {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	// Open the zip file
	file, err := os.OpenFile(archive, flags, 0644)
	if err != nil {
		return nil
	}
	defer file.Close()
	// Create zip.Writer that implements a zip file writer
	zw := zip.NewWriter(file)
	defer zw.Close()
	// Iterate through the files writing each into the zip file
	for _, filename := range files {
		// Write the file into the zip
		if err := addToArchive(filename, zw); err != nil {
			return err
		}
	}
	return nil
}

// readArchive read the contents from the zip file
func readArchive(archive string) error {
	// Open the zip file by name
	rc, err := zip.OpenReader(archive)
	if err != nil {
		return nil
	}
	defer rc.Close()
	// Iterate through the files in the zip file to read the file contents
	for _, file := range rc.File {
		frc, err := file.Open()
		if err != nil {
			return err
		}
		defer frc.Close()
		// Write the file contents to stdout
		fmt.Fprintf(os.Stdout, "Contents of the file %s:\n", file.Name)
		copied, err := io.Copy(os.Stdout, frc)
		if err != nil {
			return err
		}
		// Check the size of the file
		if uint64(copied) != file.UncompressedSize64 {
			return fmt.Errorf("length of the file contents doesn't match with the file %s", file.Name)
		}
		fmt.Println()
	}
	return nil
}

func main() {
	// Name of the zip file
	archive := "sources.zip"
	// Files to the archived
	files := []string{"main.go", "readme.txt"}
	// Archive files into zip format
	err := archiveFiles(files, archive)
	if err != nil {
		log.Fatalf("error while writing to zip file:%s\n", err)
	}
	// Read the file contents of zip file
	err = readArchive(archive)
	if err != nil {
		log.Fatalf("error while reading the zip file: %s\n", err)
	}
}
