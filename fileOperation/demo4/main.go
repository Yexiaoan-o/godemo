package main
import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func main(){
	// Open a file
	filePath := "C:/GoProjects/testFile/test.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open the file:", err)
		return
	}

	// Close the file with defer
	defer file.Close()

	// Read and output content
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// Output to the terminal
		fmt.Print(str)
	}

	
	str := "Hi Hangzhou\n"
	// Use *Writer for writing 
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++{
		writer.WriteString(str)
	}

	writer.Flush()
}