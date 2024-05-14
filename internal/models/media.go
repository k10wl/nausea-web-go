package models

type Image struct {
	URL          string `firestore:"URL"`
	ThumbnailURL string `firestore:"thumbnailURL"`
	MediaSize
}

type MediaSize struct {
	Width  int `firestore:"width"`
	Height int `firestore:"height"`
}
