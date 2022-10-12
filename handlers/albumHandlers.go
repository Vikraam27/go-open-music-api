package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Vikraam27/go-open-music-api/exceptions"
	"github.com/Vikraam27/go-open-music-api/models"
	"github.com/Vikraam27/go-open-music-api/services"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func CreateAlbumHandler(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-type", "application/json")
	var album models.AlbumPayload

	err := json.NewDecoder(req.Body).Decode(&album)
	if err != nil {
		return exceptions.NewHTTPError(err, 400, "Bad request, invalid JSON.")
	}
	validate := validator.New()
	if err := validate.Struct(album); err != nil {
		return exceptions.NewHTTPError(err, 400, "Bad request: name and year is required.")
	}
	albumId, err := services.CreateAlbumService(album)
	if err != nil {
		return err
	}

	response := models.CreateAlbumResponse{
		Status:  "success",
		Message: "successfully create album",
		Data: models.AlbumId{
			AlbumId: albumId,
		},
	}
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)
	return nil
}

func GetAlbumDetailHandler(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-type", "application/json")
	params := mux.Vars(req)

	albumDetails, err := services.GetAlbumDetailService(params["id"])

	if err != nil {
		return err
	}

	songs, err := services.GetSongByAlbumId(params["id"])

	if err != nil {
		return err
	}

	response := models.GetAlbumDetailResponse{
		Status: "success",
		Data: models.GetAlbumDetailData{
			Album: models.Album{
				ID:   albumDetails.ID,
				Name: albumDetails.Name,
				Year: albumDetails.Year,
				Song: songs,
			},
		},
	}

	json.NewEncoder(res).Encode(response)
	return nil
}

func UpdateAlbumHandler(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-type", "application/json")
	params := mux.Vars(req)

	var album models.AlbumPayload

	err := json.NewDecoder(req.Body).Decode(&album)
	if err != nil {
		return exceptions.NewHTTPError(err, 400, "Bad request, invalid JSON.")
	}
	validate := validator.New()
	if err := validate.Struct(album); err != nil {
		return exceptions.NewHTTPError(err, 400, "Bad request: name and year is required.")
	}

	rowsAffected, err := services.UpdateAlbumService(params["id"], album)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("successfully update the album, row affected %d", rowsAffected)
	response := models.ResponseWithOutData{
		Status:  "success",
		Message: msg,
	}

	json.NewEncoder(res).Encode(response)
	return nil
}

func DeleteAlbumHandler(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-type", "application/json")
	params := mux.Vars(req)

	rowsAffected, err := services.DeleteAlbumService(params["id"])
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("successfully delete the album, row affected %d", rowsAffected)
	response := models.ResponseWithOutData{
		Status:  "success",
		Message: msg,
	}

	json.NewEncoder(res).Encode(response)
	return nil
}
