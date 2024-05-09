package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Firestore struct {
	client *firestore.Client
}

func NewFirestoreClient(projectID string) *Firestore {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	return &Firestore{client: client}
}

func (f *Firestore) dataCollection() *firestore.CollectionRef {
	return f.client.Collection("data")
}

func (f *Firestore) aboutDoc() *firestore.DocumentRef {
	return f.dataCollection().Doc("about")
}

func (f *Firestore) metaDoc() *firestore.DocumentRef {
	return f.dataCollection().Doc("meta")
}
