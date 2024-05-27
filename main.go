package main

import (
	"os"
	"sync"
)

// I want to design a safe counter which will
// be used by multiple goroutines.
// 10 different goroutines will increment the counter a single time.
// The counter should be 10 at the end of the program.
type SafeCounter struct {
	mutex sync.Mutex
	data  map[string]int
}

func (counter *SafeCounter) Increment(key string, incrementNum int) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.data[key] += 1
}

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

func main() {
	chunk1 := CreateChunk("abc", "./test1.data")

	chunk1.RecordAppend([]byte("Hello"))
	chunk1.RecordAppend([]byte(" World"))
	// I want to create a chunk and add data into it in append mode twice
	// err := ioutil.WriteFile("./test.data", mydata, 0777)

	// if err != nil {
	// 	fmt.Println((err))
	// }

	// fmt.Println("Successfully wrote file")
	// // count := SafeCounter{data: make(map[string]int)}

	// file, err := os.OpenFile("./test.data", os.O_RDWR|os.O_CREATE, 0666)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// file.Seek(-1, 0)

	// file.Write([]byte("Yo"))
	// for i := 0; i < 1000; i++ {
	// 	waitGroup.Add(1)
	// 	go func() {
	// 		defer waitGroup.Done()
	// 		count.Increment("Test", 1)
	// 	}()
	// }

	// waitGroup.Wait()
	// fmt.Println(count.data)

	// test := master.NewMaster()
	// test.Add("key", 1)

	// fmt.Println(test.Get("key"))
}
