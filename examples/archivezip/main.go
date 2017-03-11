package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

// addToArchive writes a given file into a zip file.
func addToArchive(filename string, zw *zip.Writer) error {
	// Open the given file to archive into a zip file.
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// Create adds a file to the zip file using the given name/
	// Create returns a io.Writer to which the file contents should be written.
	wr, err := zw.Create(filename)
	if err != nil {
		return err
	}
	// Write the file contents to the zip file.
	if _, err := io.Copy(wr, file); err != nil {
		return err
	}
	return nil
}

// archiveFiles archives a group of given files into a zip file.
func archiveFiles(files []string, archive string) error {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	// Open the tar file
	file, err := os.OpenFile(archive, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// Create zip.Writer that implements a zip file writer.
	zw := zip.NewWriter(file)
	defer zw.Close()
	// Iterate through the files to write each file into the zip file.
	for _, filename := range files {
		// Write the file into tar file.
		if err := addToArchive(filename, zw); err != nil {
			return err
		}
	}
	return nil
}

// readArchive reads the file contents from tar file.
func readArchive(archive string) error {
	// Open the zip file specified by name and return a ReadCloser.
	rc, err := zip.OpenReader(archive)	
	if err != nil {
		return err
	}
	defer rc.Close()
	// Iterate through the files in the zip file to read the file contents.
	for _, file := range rc.File {
		frc, err := file.Open()
		if err != nil {
			return err
		}
		defer frc.Close()
		fmt.Fprintf(os.Stdout, "Contents of the file %s:\n", file.Name)
		// Write the contents into Stdout
		copied, err := io.Copy(os.Stdout, frc)
		if err != nil {
			return err
		}
		// Check the size of the file.
		if uint64(copied) != file.UncompressedSize64 {
			return fmt.Errorf("Length of the file contents doesn't match with the file %s", file.Name)
		}
		fmt.Println()
	}
	return nil
}

func main() {
	// Name of the zip file
	archive := "source.zip"
	// Files to be archived in zip format.
	files := []string{"main.go", "readme.txt"}
	// Archive files into zip format.
	err := archiveFiles(files, archive)
	if err != nil {
		log.Fatalf("Error while writing to zip file:%s\n", err)
	}
	// Read the file contents of tar file.
	err = readArchive(archive)
	if err != nil {
		log.Fatalf("Error while reading the zip file:%s\n", err)

	}
}
