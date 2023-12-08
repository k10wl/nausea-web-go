package main

type Store struct {
	firebase FirebaseClient
}

type About struct {
	Info string
}

func NewStore(firebase FirebaseClient) *Store {
	firebase.GetInfo()
	return &Store{
		firebase: firebase,
	}
}

func (s *Store) GetAbout() Info {
	return s.firebase.GetInfo()
}
