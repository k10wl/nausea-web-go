package firestore

import (
	"context"
	"nausea-web/internal/models"
)

func (f *Firestore) GetAbout(ctx context.Context) (*models.About, error) {
	snapshot, err := f.aboutDoc().Get(ctx)
	if err != nil {
		return nil, err
	}
	var about *models.About
	snapshot.DataTo(&about)
	return about, nil
}

func (f *Firestore) GetContacts(ctx context.Context) (*models.Contacts, error) {
	snapshot, err := f.contactsDoc().Get(ctx)
	if err != nil {
		return nil, err
	}
	var contacts *models.Contacts
	snapshot.DataTo(&contacts)
	return contacts, nil
}

func (f *Firestore) GetMeta(ctx context.Context) (*models.Meta, error) {
	snapshot, err := f.metaDoc().Get(ctx)
	if err != nil {
		return nil, err
	}
	var meta *models.Meta
	snapshot.DataTo(&meta)
	return meta, nil
}
