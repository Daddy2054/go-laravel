package celeritas

import (
	"fmt"
	"github.com/daddy2054/celeritas/filesystems"
	"io"
	"net/http"
	"os"
	"path"
)

func (c *Celeritas) UploadFile(r *http.Request, destination, field string, fs filesystems.FS) error {
	fileName, err := c.getFileToUpload(r, field)
	if err != nil {
		c.ErrorLog.Println(err)
		return err
	}

	if fs != nil {
		err = fs.Put(fileName, destination)
		if err != nil {
			c.ErrorLog.Println(err)
			return err
		}
	} else {
		err = os.Rename(fileName, fmt.Sprintf("%s/%s", destination, path.Base(fileName)))
		if err != nil {
			c.ErrorLog.Println(err)
			return err
		}
	}
	return nil
}

func (c *Celeritas) getFileToUpload(r *http.Request, fieldName string) (string, error) {
	_ = r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile(fieldName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	dst, err := os.Create(fmt.Sprintf("./tmp/%s", header.Filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("./tmp/%s", header.Filename), nil
}