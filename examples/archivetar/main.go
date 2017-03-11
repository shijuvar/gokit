package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

// addToArchive writes a given file into a .tar file
// Returns nill if the operation is succeeded
func addToArchive(filename string, tw *tar.Writer) error {
	// Open the file to archive into tar file.
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// Get the FileInfo struct that describe the file.
	fileinfo, err := file.Stat()
	// Create a pointer to tar.Header struct
	hdr := &tar.Header{
		ModTime: fileinfo.ModTime(),            // modified time
		Name:    filename,                      // name of header
		Size:    fileinfo.Size(),               // length in bytes
		Mode:    int64(fileinfo.Mode().Perm()), // permission and mode bits
	}
	// WriteHeader writes tar.Header and prepares to accept the file's contents.
	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	// Write the file contents to the tar file.
	copied, err := io.Copy(tw, file)
	if err != nil {
		return err
	}
	// Check the size of copied file with the source file.
	if copied < fileinfo.Size() {
		return fmt.Errorf("Size of the copied file doesn't match with source file %s: %s", filename, err)
	}
	return nil
}

// archiveFiles archives a group of given files into a tar file.
func archiveFiles(files []string, archive string) error {
	// Flags for open the tar file.
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	// Open the tar file
	file, err := os.OpenFile(archive, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// Create a new Writer writing to given file object.
	// Writer provides sequential writing of a tar archive in POSIX.1 format.
	tw := tar.NewWriter(file)
	defer tw.Close()
	// Iterate through the files to write each file into the tar file.
	for _, filename := range files {
		// Write the file into tar file.
		if err := addToArchive(filename, tw); err != nil {
			return err
		}
	}
	return nil
}

// readArchive reads the file contents from tar file.
func readArchive(archive string) error {
	// Open the tar archive file.
	file, err := os.Open(archive)
	if err != nil {
		return err
	}
	defer file.Close()
	// Create the tar.Reader to read the tar archive.
	// A Reader provides sequential access to the contents of a tar archive.
	tr := tar.NewReader(file)
	// Iterate through the files in the tar archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			fmt.Println("end")
			break
		}
		if err != nil {
			return err
		}
		size := hdr.Size
		contents := make([]byte, size)
		read, err := io.ReadFull(tr, contents)
		// Check the size of file contents
		if int64(read) != size {
			return fmt.Errorf("Size of the opened file doesn't match with the file %s", hdr.Name)
		}
		// hdr.Name returns the file name.
		fmt.Printf("Contents of the file %s:\n", hdr.Name)
		// Writing the file contents into Stdout.
		fmt.Fprintf(os.Stdout, "\n%s", contents)
	}
	return nil
}

func main() {
	// Name of the tar file
	archive := "source.tar"
	// Files to be archived in tar format
	files := []string{"main.go", "readme.txt"}
	// Archive files into tar format
	err := archiveFiles(files, archive)
	if err != nil {
		log.Fatalf("Error while writing to tar file:%s", err)
	}
	// Archiving is sucsess.
	fmt.Println("The tar file source.tar has been created")
	// Read the file contents of tar file
	err = readArchive(archive)
	if err != nil {
		log.Fatalf("Error while reading the tar file:%s", err)
	}
}
