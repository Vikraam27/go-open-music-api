package models

type ResponseWithOutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type AlbumPayload struct {
	Name string `json:"name" validate:"required"`
	Year uint16 `json:"year" validate:"required"`
}
type Album struct {
	ID   string       `json:"id"`
	Name string       `json:"name"`
	Year int          `json:"year"`
	Song []MappedSong `json:"songs"`
}

type CreateAlbumResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    AlbumId `json:"data"`
}

type GetAlbumDetailResponse struct {
	Status string             `json:"status"`
	Data   GetAlbumDetailData `json:"data"`
}

type GetAlbumDetailData struct {
	Album Album `json:"album"`
}

type AlbumId struct {
	AlbumId string `json:"albumId"`
}
