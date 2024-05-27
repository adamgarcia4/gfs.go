package chunk

import (
	"log"
	"os"
)

const MAX_FILE_SIZE = 4096

type Chunk struct {
	chunkId string
	path    string
	filePtr *os.File
}

func CreateChunk(chunkId string, path string) Chunk {
	chunk := Chunk{}
	chunk.chunkId = chunkId
	chunk.path = path
	chunk.filePtr, _ = os.OpenFile(chunk.path, os.O_RDWR|os.O_CREATE, 0666)

	return chunk
}

func (chunk *Chunk) RecordAppend(dataToAppend []byte) {
	chunk.filePtr.Seek(0, 2)
	chunk.filePtr.Write(dataToAppend)
}

func (chunk *Chunk) GetSize() int64 {
	test, err := chunk.filePtr.Stat()

	if err != nil {
		log.Fatal(err)
		return -1
	}

	return test.Size()
}
