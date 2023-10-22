package tutorial

import (
	"bufio"
	_ "embed"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err)
	}
}

func pathTutorial() {
	log.Println("File paths usage")
	// Path validation and usage
	p1 := filepath.Join("parent_dir", "sub_dir", "filename.yaml")
	p2 := filepath.Join("parent_dir/sub_dir", "filename.yaml")
	log.Println(p1)
	log.Println(p2)

	parentDir := filepath.Dir(p2)
	fileBase := filepath.Base(p2)
	fileExt := filepath.Ext(p2)
	fileName := strings.TrimSuffix(fileBase, ".yaml")
	log.Println(parentDir)
	log.Println(fileBase)
	log.Println(fileExt)
	log.Println(fileName)
}

func writeTutorial() {
	log.Println("Write in files usage")
	// dump a string (or just bytes) into a file
	d1 := []byte("hello\nworld")
	err := os.WriteFile("/tmp/writeTutorial1.txt", d1, 0600)
	fatalIfErr(err)
	log.Println("Check for '/tmp/writeTutorial1.txt' content")

	// For more granular writes, open a file for writing
	// It’s idiomatic to defer a Close immediately after opening a file.
	f, err := os.Create("/tmp/writeTutorial2.txt")
	fatalIfErr(err)
	// closing file at the end of this method execution
	// you can also call f.Close() at any moment in this method
	defer f.Close()
	log.Printf("Check for '%s' content", f.Name())

	// Writing strings
	// Easier for simple string but less convenient than buffered writer for
	// large content
	_, err = f.WriteString("hello world !\nhello again !\n")
	fatalIfErr(err)
	// should sync after strings writing
	// Issue a Sync to flush writes to stable storage
	f.Sync()

	// Buffered writing (better for big content).
	// The advantage of using BufferedReader or BufferedWriter is that it
	// reduces the number of physical reads from and writes to the disk. It
	// makes use of a buffering system and performs reading/writing all at once.
	// Hence, there is more efficiency. You can try reading/writing a large file
	// with and without using BufferedReader/BufferedWriter and see the
	// difference.
	w := bufio.NewWriter(f)
	_, err = w.WriteString("Very large content is better with buffered writer")
	fatalIfErr(err)
	// Use Flush to ensure all buffered operations have been applied to the
	// underlying writer.
	w.Flush()

	writeLargeFile("/tmp/testWrite.txt")
}

func simpleWrite(path string, content []byte) {
	err := os.WriteFile(path, content, 0660)
	fatalIfErr(err)
}

func writeLargeFile(path string) {
	// write a large file
	f2, err := os.Create(path)
	defer f2.Close()
	fatalIfErr(err)
	w2 := bufio.NewWriter(f2)

	log.Printf("Writing a large file in %s", path)
	for i := 0; i <= 1000; i++ {
		_, err = w2.WriteString("test ")
		fatalIfErr(err)
		w2.Flush()
	}
}

func readTutorial() {
	log.Println("Read files usage")
	filePath := "/tmp/golang_starter_read_tutorial.txt"
	// writing for further usage
	simpleWrite(filePath, []byte("hello world !"))

	// Reading files requires checking most calls for errors.
	contentBytes, err := os.ReadFile(filePath)
	fatalIfErr(err)
	// must convert output
	contentString := string(contentBytes)
	log.Println(contentString) // should read 'hello world !'

	// For more granular readings, open a file for further usage
	// It’s idiomatic to defer a Close immediately after opening a file.
	f, err := os.Open(filePath)
	fatalIfErr(err)

	// read a certain amount of bytes
	myBuffer := make([]byte, 256)
	// f.Read() return the number of bytes that were actually read in the file
	byteRead, err := f.Read(myBuffer)
	fatalIfErr(err)
	log.Printf("Have read %d bytes in file", byteRead)
	log.Println(string(myBuffer[:byteRead]))
	// You can close the file now
	f.Close()
	// !!! For further readings, you can use the function:
	// Seek(0, 0)
	// to set the offset for reading at the beginning of the file
	// Check this behavior further !

	// You can also read less bytes than the file contains with a small
	// buffer, so be careful
	mySmallBuffer := make([]byte, 5)
	// re-opening the file for reading
	f, err = os.Open(filePath)
	fatalIfErr(err)
	_, err = f.Read(mySmallBuffer)
	fatalIfErr(err)
	log.Println("Reading less bytes than there are in the file :")
	log.Println(string(mySmallBuffer))
	// NOT CLOSING THE FILE BECAUSE WE ARE USING THE Seek() function

	// Read from an offset in file
	b1 := make([]byte, 5)
	// Seek sets the offset for the next Read or Write on file to offset,
	// interpreted according to whence: 0 means relative to the origin of the
	// file, 1 means relative to the current offset, and 2 means relative to
	// the end. It returns the new offset and an error, if any.
	o1, err := f.Seek(6, 0)
	fatalIfErr(err)
	// The offset is set to 6th byte in the file
	// So we expect to read after the 'hello' word
	c1, err := f.Read(b1)
	fatalIfErr(err)
	log.Printf("Read %d bytes from position %d in file:", c1, o1)
	log.Println(string(b1))

	// reset reading
	_, err = f.Seek(0, 0)
	fatalIfErr(err)

	// For large readings, use a buffered reader.
	// Indeed, if you have a 32GB file, you will hardly load it into the memory
	// using the methods above.
	// The idea is to put into a buffer a part of a file to read it and
	// continue to read until the end of the file.
	// This method is the safest because you fully control the amount of bytes
	// you read at each iteration.
	filePath2 := "/tmp/testReadLarge.txt"
	writeLargeFile(filePath2)
	f2, err := os.Open(filePath2)
	fatalIfErr(err)
	// starting an infinite loop breaking at the end of a file
	log.Println("Reading a large file")
	stats, err := f2.Stat()
	fatalIfErr(err)
	log.Printf("Large file size: %d bytes", stats.Size())
	// reading
	br := bufio.NewReader(f2)
	totalBytesRead := 0
	for {
		// reading 1024 bytes at each iteration
		largeBuffer := make([]byte, 1024)
		bytesRead, err := br.Read(largeBuffer)
		// checking if bytes were read
		if bytesRead == 0 {
			if err == io.EOF {
				log.Printf("End of file %s", filePath2)
				break
			}
			fatalIfErr(err)
		}
		log.Printf("Read %d in file %s", bytesRead, filePath2)
		largeBuffer = largeBuffer[:bytesRead]
		totalBytesRead += bytesRead
	}
	log.Println("Finished reading")
	log.Printf("Read %d bytes in total", totalBytesRead)

	_, err = f2.Seek(0, 0)
	fatalIfErr(err)

	// For large readings, you can also use a Scanner.
	// A Scanner is convenient for large files while being simple in use.
	// A Scanner uses Tokens to store parts of a file it reads.
	// However, there is a risk of flooding the memory according to the type of
	// token you read.
	// For example, if you use Text(), the token is a line, but if the line is
	// too long, you can end up with an error.
	log.Println("Read large file using a scanner")
	scanner := bufio.NewScanner(f2)
	totalLinesRead := 0
	for scanner.Scan() {
		// read each line
		line := scanner.Text()
		totalLinesRead++
		log.Println(line)
	}

	f.Close()
}

func Files() {
	//pathTutorial()
	//writeTutorial()
	readTutorial()
}
