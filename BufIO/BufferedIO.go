package main

import (
	"bufio"
	"fmt"
	"os"
)

// Buffered I/O is an efficient way to handle files.
// It is faster to write a big chunk of data than many small chunks.
// The bufio-package contains many more useful functions. Explore it!

// Writing a textfile
func WriteFile(filename string) {
	f, err := os.Create(filename) // creating the textfile
	if err != nil {
		panic(err.Error()) // error handling
	}
	defer f.Close()
	w := bufio.NewWriter(f) // new writer with default size (4096 Bytes)
	var text string
	for i := 0; i < 100; i++ {
		text = fmt.Sprintf("%s %02d\t", "Message", i)
		w.WriteString(text) // write to the buffer
	}
	w.Flush() //write everything to the underlying writer. Never forget this!
}

// Reading a textfile
func ReadFile(filename string) {
	f, err := os.Open(filename) // open the textfile
	if err != nil {
		panic(err.Error()) // error handling
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error()) // error handling
	}
}

func main() {
	const filename = "Strings.txt"
	WriteFile(filename)
	fmt.Println("100 useless strings written to the file...")
	fmt.Println("Now we will read the file.")
	fmt.Println()
	ReadFile(filename)
}
