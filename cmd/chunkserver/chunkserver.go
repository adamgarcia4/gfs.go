package main

import (
	"fmt"

	chunk "github.com/adamgarcia4/gfs.go/packages/chunk"
)

func main() {
	fmt.Println("Chunkserver starting")
	chunk1 := chunk.CreateChunk("abc", "./test1.data")

	chunk1.RecordAppend([]byte("Hello"))
	chunk1.RecordAppend([]byte(" World"))
}
