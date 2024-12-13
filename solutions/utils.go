package solutions

import (
	"fmt"
	"os"
)

func WriteToFile(content string, filename string) {
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Println("Failed to create file: ", filename)
		fmt.Println("err1: ", err1)
	}
	defer f.Close()

	nbytes, err2 := f.WriteString(content)
	if err2 != nil {
		fmt.Println("Failed to write to file: ", filename)
		fmt.Println("err2: ", err2)
	}
	fmt.Println("Wrote ", nbytes, " bytes to file: ", filename)
}
