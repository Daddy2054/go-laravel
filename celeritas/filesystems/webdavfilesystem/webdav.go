package webdavfilesystem

import (
	"fmt"
	"os"
	"path"

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

func (s *WebDAV) List(prefix string) ([]filesystems.Listing, error) {
	var listing []filesystems.Listing
	return listing, nil
}

func (s *WebDAV) Delete(itemsToDelete []string) bool {
	return true
}

func (s *WebDAV) Get(destination string, items ...string) error {
	return nil
}
