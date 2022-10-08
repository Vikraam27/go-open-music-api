package models

type SongPayload struct {
	Title     string
	Year      uint16
	Genre     string
	Performer string
	Duration  uint16
	AlbumId   string
}
