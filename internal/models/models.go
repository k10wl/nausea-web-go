package models

import "time"

type ID struct {
	ID string `firestore:"id"`
}

type Timestamps struct {
	CreatedAt time.Time  `firestore:"createdAt"`
	UpdatedAt time.Time  `firestore:"updatedAt"`
	DeletedAt *time.Time `firestore:"deletedAt"`
}
