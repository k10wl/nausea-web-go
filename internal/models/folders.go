package models

type Folder struct {
	ID
	Timestamps
	ParentID       string          `firestore:"parentID"`
	Name           string          `firestore:"name"`
	FolderContents []FolderContent `firestore:"folders"`
	MediaContents  []MediaContent  `firestore:"media"`
	Protected      bool            `firestore:"protected"`
	ProhibitNested bool            `firestore:"prohibitNested"`
	ProhibitMedia  bool            `firestore:"prohibitMedia"`
}

type ContentBase struct {
	ID
	Timestamps
	RefID string `firestore:"refID"`
}

type FolderContent struct {
	ContentBase
	Name string `firestore:"name"`
}

type MediaContent struct {
	ContentBase
	MediaSize
	URL          string `firestore:"URL"`
	ThumbnailURL string `firestore:"thumbnailURL"`
	Name         string `firestore:"name"`
	ParentID     string `firestore:"parentID"`
	Description  string `firestore:"description"`
}
