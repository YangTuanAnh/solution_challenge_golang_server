package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"log"
	"SC2024/pkg/firebase"
)

var (
	defaultBucket *storage.BucketHandle
)

type FirebaseStorage struct {
}

func Setup(ctx context.Context) error {
	client, err := firebase.GetFirebaseApp().Storage(ctx)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket("hcmuwus-rehabox.appspot.com")
	if err != nil {
		return err
	}

	defaultBucket = bucket

	return nil
}

func DefaultBucket() *storage.BucketHandle {
	return defaultBucket
}

func (fs *FirebaseStorage) UploadSingle() {

	attr, _ := defaultBucket.Object("atilla.txt").Attrs(context.Background())
	log.Println(attr.MediaLink)
	for key, val := range attr.Metadata {
		log.Println(key, val)

	}

}