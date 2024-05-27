package main

import (
	"fmt"
	"os"

	chunk "github.com/adamgarcia4/gfs.go/packages/chunk"
)

func main() {
	fmt.Println("Chunkserver starting")
	chunk1 := chunk.CreateChunk("abc", "./test1.data")

	size := chunk1.GetSize()

	fmt.Println("Size is: ", size)

	file, _ := os.ReadFile("./10b.txt")

	chunk1.RecordAppend([]byte(file))
	// chunk1.RecordAppend([]byte(" World"))
}
