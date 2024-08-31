package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	// Open a file
	filePath := "C:/GoProjects/testFile/test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Failed to open the file:", err)
		return
	}

	// Close the file with defer
	defer file.Close()

	// Write "Hello, sean" for 5 times
	str := "NIHAO sean\n"
	// Use *Writer for writing 
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++{
		writer.WriteString(str)
	}

	writer.Flush()
}