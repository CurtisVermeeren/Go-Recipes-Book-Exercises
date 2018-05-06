package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

// addToArchive writes a given file into a tar file
func addToArchive(filename string, tw *tar.Writer) error {

	// Open the file to archive into the tar
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get the FileInfo struct that describes the file
	fileinfo, err := file.Stat()

	// Create a pointer to tar.Header struct
	header := &tar.Header{
		ModTime: fileinfo.ModTime(),
		Name:    fileinfo.Name(),
		Size:    fileinfo.Size(),
		Mode:    int64(fileinfo.Mode().Perm()),
	}

	// Write the header for the tar file
	if err := tw.WriteHeader(header); err != nil {
		return err
	}

	// Write the contents of the tar file
	copied, err := io.Copy(tw, file)
	if err != nil {
		return err
	}

	// Check that the number of bytes copied from the source is correct
	if copied < fileinfo.Size() {
		return fmt.Errorf("Size of the copied file doesn't match with source file %s: %s", filename, err)
	}

	return nil
}

// archiveFiles writes a group of files into a tar file
func archiveFiles(files []string, archive string) error {
	// Flags for opening the tar file
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC

	// Open the tar file
	file, err := os.OpenFile(archive, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new writer for writing to a given file object
	// Writer provides sequential writing of a tar archive in POSIX.1 format
	tw := tar.NewWriter(file)
	defer tw.Close()

	// Iterate through the files and write them into the tar file
	for _, filename := range files {
		// Write the file into the tar
		if err := addToArchive(filename, tw); err != nil {
			return err
		}
	}

	return nil
}

// readArchive reads the file content from a tar archive file
func readArchive(archive string) error {
	// Open the tar file
	file, err := os.Open(archive)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a tar.Reader to read the archive
	// Reader provides sequential access to the contents of a tar archive
	tr := tar.NewReader(file)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}

		if err != nil {
			return err
		}

		size := header.Size
		contents := make([]byte, size)
		read, err := io.ReadFull(tr, contents)
		// Check the size of file contents
		if int64(read) != size {
			return fmt.Errorf("Size of the opened file deosn't match with the file %s", header.Name)
		}
		fmt.Printf("Contents of the file %s:\n", header.Name)
		// Writing the file contents into Stdout.
		fmt.Fprintf(os.Stdout, "\n%s", contents)
	}
	return nil
}

func main() {
	// Name of the tar file
	archive := "source.tar"
	// Files to be archived into a tar
	files := []string{"main.go", "readme.txt"}
	// Archive the files
	err := archiveFiles(files, archive)
	if err != nil {
		log.Fatalf("error while writing to tar file:%s", err)
	}
	// Archiving is successful.
	fmt.Println("The tar file source.tar has been created")
	// Read the file contents of tar file
	err = readArchive(archive)
	if err != nil {
		log.Fatalf("Error while reading the tar file:%s", err)
	}
}
