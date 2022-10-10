package models

type SongPayload struct {
	Title     string `json:"title" validate:"required"`
	Year      uint16 `json:"year" validate:"required"`
	Genre     string `json:"genre" validate:"required"`
	Performer string `json:"performer" validate:"required"`
	Duration  uint16 `json:"duration" validate:"required"`
}

type AddSongResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    SongId `json:"data"`
}

type SongId struct {
	SongId string `json:"songId"`
}
