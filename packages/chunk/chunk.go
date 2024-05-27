package chunk

import "os"

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
