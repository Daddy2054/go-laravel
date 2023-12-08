package filesystems

import "time"

// FS is the interface for file systems. In order to saytidfy the interface,
// all of its functions must exist
type FS interface {
	Put(filename, folder string) error
	Get(destination string, items ...string) error
	List(prefix string) ([]Listing, error)
	Delete(itemsToDelete []string) bool
}

// "Listing" desscribes one file on remote file system
type Listing struct {
	Etag         string
	LastModified time.Time
	Key          string
	Size         float64
	IsDir        bool
}
