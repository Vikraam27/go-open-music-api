package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Vikraam27/go-open-music-api/exceptions"
	"github.com/Vikraam27/go-open-music-api/models"
	"github.com/Vikraam27/go-open-music-api/services"
	"github.com/go-playground/validator/v10"
)

func AddSongHandler(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-type", "application/json")
	var song models.SongPayload

	err := json.NewDecoder(req.Body).Decode(&song)
	if err != nil {
		return exceptions.NewHTTPError(err, 400, "Bad request, invalid JSON.")
	}
	validate := validator.New()
	if err := validate.Struct(song); err != nil {
		return exceptions.NewHTTPError(err, 400, "Bad request: all property is required.")
	}
	songId, err := services.AddSongService(song)

	if err != nil {
		return err
	}

	response := models.AddSongResponse{
		Status:  "success",
		Message: "successfulu add song",
		Data: models.SongId{
			SongId: songId,
		},
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)
	return nil
}
