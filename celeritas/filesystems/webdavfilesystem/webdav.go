package webdavfilesystem

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/daddy2054/celeritas/filesystems"
	"github.com/studio-b12/gowebdav"
)

type WebDAV struct {
	Host string
	User string
	Pass string
}

func (w *WebDAV) getCredentials() *gowebdav.Client {
	c := gowebdav.NewClient(w.Host, w.User, w.Pass)
	return c
}

func (w *WebDAV) Put(fileName, folder string) error {
	client := w.getCredentials()

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	err = client.WriteStream(fmt.Sprintf("%s/%s", folder, path.Base(fileName)), file, 0664)
	if err != nil {
		return err
	}
	return nil
}

// List lists files on the remote file system
func (w *WebDAV) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing
	client := w.getCredentials()
	files, err := client.ReadDir(prefix)
	if err != nil {
		return listing, err
	}

	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			b := float64(file.Size())
			kb := b / 1024
			mb := kb / 1024
			current := filesystems.Listing{
				LastModified: file.ModTime(),
				Key:          file.Name(),
				Size:         mb,
				IsDir:        file.IsDir(),
			}
			listing = append(listing, current)
		}
	}

	return listing, nil
}

// Delete deletes file(s) from the remote file system
func (w *WebDAV) Delete(itemsToDelete []string) bool {
	client := w.getCredentials()
	for _, item := range itemsToDelete {
		err := client.Remove(item)
		if err != nil {
			return false
		}
	}
	return true
}

func (s *WebDAV) Get(destination string, items ...string) error {
	return nil
}
