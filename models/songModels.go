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

type MappedSong struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Performer string `json:"performer"`
}

type GetAllSongsResponse struct {
	Status string     `json:"status"`
	Data   *SongsData `json:"data"`
}

type SongsData struct {
	Songs []MappedSong `json:"songs"`
}

type Song struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Year      uint16 `json:"year"`
	Genre     string `json:"genre"`
	Performer string `json:"performer"`
	Duration  uint16 `json:"duration"`
}

type GetSongByIdResponse struct {
	Status string   `json:"status"`
	Data   SongData `json:"data"`
}

type SongData struct {
	Songs Song `json:"song"`
}
