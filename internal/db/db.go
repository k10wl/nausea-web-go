package db

import (
	"context"
	"nausea-web/internal/models"
	"time"
)

type IDB interface {
	GetAbout(ctx context.Context) (*models.About, error)
	GetContacts(ctx context.Context) (*models.Contacts, error)
	GetMeta(ctx context.Context) (*models.Meta, error)
	GetFolder(ctx context.Context, id string) (*models.Folder, error)
}

type DB struct {
	metaCache *models.Meta
	client    IDB
}

func NewDB(client IDB) *DB {
	return &DB{client: client}
}

func (db *DB) GetAbout(ctx context.Context) (*models.About, error) {
	return db.client.GetAbout(ctx)
}

func (db *DB) GetContacts(ctx context.Context) (*models.Contacts, error) {
	return db.client.GetContacts(ctx)
}

func (db *DB) GetMeta(ctx context.Context) (*models.Meta, error) {
	if db.metaCache != nil {
		return db.metaCache, nil
	}
	go func() {
		time.Sleep(5 * time.Minute)
		db.metaCache = nil
	}()
	meta, err := db.client.GetMeta(ctx)
	if err != nil {
		return nil, err
	}
	db.metaCache = meta
	return meta, nil
}

func (db *DB) GetFolder(ctx context.Context, id string) (*models.Folder, error) {
	return db.client.GetFolder(ctx, id)
}
