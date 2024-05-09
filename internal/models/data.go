package models

type About struct {
	Bio   string `firestore:"bio"`
	Image Image  `firestore:"image"`
}

type Contacts struct {
	Links string `firestore:"links"`
}

type Meta struct {
	Background Image `firestore:"background"`
}
