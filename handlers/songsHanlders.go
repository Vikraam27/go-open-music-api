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
	return err
}

func GetAllSongsHanlder(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-type", "application/json")

	query := req.URL.Query()

	if query.Get("title") != "" {
		fmt.Println("excuted")
		songs, err := services.GetAllSongsByTitleService(query.Get("title"))
		if err != nil {
			return err
		}

		response := models.GetAllSongsResponse{
			Status: "success",
			Data: &models.SongsData{
				Songs: songs,
			},
		}
		json.NewEncoder(res).Encode(response)
		return err
	} else if query.Get("performer") != "" {
		songs, err := services.GetAllSongsByPerformerService(query.Get("performer"))
		if err != nil {
			return err
		}

		response := models.GetAllSongsResponse{
			Status: "success",
			Data: &models.SongsData{
				Songs: songs,
			},
		}
		json.NewEncoder(res).Encode(response)
		return err
	} else if (query.Get("title") != "") && (query.Get("performer") != "") {
		songs, err := services.GetAllSongsByTitleAndPerformerService(query.Get("title"), query.Get("performer"))
		if err != nil {
			return err
		}

		response := models.GetAllSongsResponse{
			Status: "success",
			Data: &models.SongsData{
				Songs: songs,
			},
		}
		json.NewEncoder(res).Encode(response)
		return err
	} else {
		songs, err := services.GetAllSongsService()
		if err != nil {
			return err
		}

		response := models.GetAllSongsResponse{
			Status: "success",
			Data: &models.SongsData{
				Songs: songs,
			},
		}
		json.NewEncoder(res).Encode(response)
		return err

	}
}

func GetSongByIdHandler(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-type", "application/json")
	params := mux.Vars(req)
	song, err := services.GetSongByIdService(params["id"])

	if err != nil {
		return err
	}

	response := models.GetSongByIdResponse{
		Status: "success",
		Data: models.SongData{
			Songs: song,
		},
	}

	json.NewEncoder(res).Encode(response)
	return err
}

func UpdateSongHandler(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-type", "application/json")
	params := mux.Vars(req)

	var song models.SongPayload

	err := json.NewDecoder(req.Body).Decode(&song)
	if err != nil {
		return exceptions.NewHTTPError(err, 400, "Bad request, invalid JSON.")
	}
	validate := validator.New()
	if err := validate.Struct(song); err != nil {
		return exceptions.NewHTTPError(err, 400, "Bad request: all property is required.")
	}

	rowsAffected, err := services.UpdateSongService(params["id"], song)

	if err != nil {
		return err
	}
	msg := fmt.Sprintf("successfully update song, row affected %d", rowsAffected)
	response := models.ResponseWithOutData{
		Status:  "success",
		Message: msg,
	}

	json.NewEncoder(res).Encode(response)
	return nil
}

func DeleteSongHandler(res http.ResponseWriter, req *http.Request) error {
	res.Header().Add("Content-type", "application/json")
	params := mux.Vars(req)

	rowsAffected, err := services.DeleteSongService(params["id"])

	if err != nil {
		return err
	}
	msg := fmt.Sprintf("successfully delete song, row affected %d", rowsAffected)
	response := models.ResponseWithOutData{
		Status:  "success",
		Message: msg,
	}

	json.NewEncoder(res).Encode(response)
	return nil
}
