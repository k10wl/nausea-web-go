package firestore

import (
	"context"
	"nausea-web/internal/models"
)

func (f *Firestore) GetFolder(ctx context.Context, id string) (*models.Folder, error) {
	var folder models.Folder
	snap, err := f.foldersCollection().Doc(id).Get(ctx)
	if err != nil {
		return &folder, err
	}
	err = snap.DataTo(&folder)
	return &folder, err
}
