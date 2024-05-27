package master

import "fmt"

// Application will ask the GFS master for the chunk handle and locations on the chunk servers.
// The master will return the chunk handle and locations to the application.

type ChunkLocations struct {
	chunkHandle string
	// Right now, I'm going to have locations be different ports that are listened to.
	locations []string
}

type Master struct {
	// internal state
	state map[string]ChunkLocations
}

func NewMaster() *Master {
	return &Master{state: make(map[string]ChunkLocations)}
}

func (m *Master) Add(fileName string, chunkIdx int) {
	// Serialize fileName and chunkIdx into a key
	fmt.Sprintf()
	key := fileName + string(chunkIdx)

	m.state[key] = ChunkLocations{chunkHandle: key, locations: []string{"8080"}}
}