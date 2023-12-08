package miniofilesystem

import (
	"context"
	"fmt"
	"log"
	"path"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	endpoint string
	Key      string
	Secret   string
	UseSSL   bool
	Region   string
	Bucket   string
}

func (m *Minio) getCredentials() *minio.Client {
	client, err := minio.New(m.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.Key, m.Secret, ""),
		Secure: m.UseSSL,
	})

	if err != nil {
		log.Println(err)
	}
	return client
}

func (m *Minio) Put(fileName, folder string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	objectName := path.Base(fileName)
	client := m.getCredentials()
	uploadInfo, err := client.FPutObject(ctx, m.Bucket, fmt.Sprintf("%s/%s", folder, objectName), fileName, minio.PutObjectOptions{})
	if err != nil {
		log.Println("Failed with FPutObject")
		log.Println(err)
		log.Println("UploadInfo:", uploadInfo)
		return err
	}
	return nil
}
