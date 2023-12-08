package main

import (
	"context"

	"cloud.google.com/go/firestore"
)

type FirebaseClient struct {
	client *firestore.Client
	ctx    context.Context
}

type Info struct {
	Bio string `firestore:"bio"`
}

func NewFirebaseClient(ctx context.Context, projectId string) *FirebaseClient {
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		panic(err)
	}
	return &FirebaseClient{
		client: client,
		ctx:    ctx,
	}
}

func (c *FirebaseClient) GetInfo() Info {
	about := c.client.Doc("about/info")
	dc, err := about.Get(c.ctx)
	var info Info
	dc.DataTo(&info)
	if err != nil {
		panic(err)
	}
	return info
}
