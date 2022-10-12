package services

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/Vikraam27/go-open-music-api/database"
	"github.com/Vikraam27/go-open-music-api/exceptions"
	"github.com/Vikraam27/go-open-music-api/models"
)

func AddSongService(song models.SongPayload) (string, error) {
	db := database.CreateConnection()
	defer db.Close()

	id := fmt.Sprintf("song-%v", strconv.Itoa(rand.Intn(9999999)))
	query := "INSERT INTO songs VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	var songId string
	err := db.QueryRow(query, id, song.Title, song.Year, song.Genre, song.Performer, song.Duration, song.AlbumId).Scan(&songId)

	if err != nil {
		return "", fmt.Errorf("fail to add song : %v", err)
	}
	return id, err
}

func GetAllSongsService() ([]models.MappedSong, error) {
	db := database.CreateConnection()
	defer db.Close()

	var songs []models.MappedSong
	songs = make([]models.MappedSong, 0, 1)

	query := "SELECT id, title, performer FROM songs"

	rows, err := db.Query(query)

	if err != nil {
		return songs, err
	}
	defer rows.Close()

	for rows.Next() {
		var song models.MappedSong

		err = rows.Scan(&song.Id, &song.Title, &song.Performer)
		if err != nil {
			return songs, err
		}
		songs = append(songs, song)
	}

	return songs, err
}

func GetSongByIdService(id string) (models.Song, error) {
	db := database.CreateConnection()
	defer db.Close()

	var song models.Song

	query := "SELECT * FROM songs WHERE id=$1"

	row := db.QueryRow(query, id)
	err := row.Scan(&song.Id, &song.Title, &song.Year, &song.Genre, &song.Performer, &song.Duration, &song.AlbumId)
	switch err {
	case sql.ErrNoRows:
		return song, exceptions.NewHTTPError(err, 404, "song not found")
	case nil:
		return song, nil
	default:
		return song, err
	}
}

func UpdateSongService(id string, song models.SongPayload) (int64, error) {
	db := database.CreateConnection()
	defer db.Close()

	query := "UPDATE songs SET title=$2, year=$3, performer=$4, genre=$5, duration=$6 WHERE id=$1"
	res, err := db.Exec(query, id, song.Title, song.Year, song.Performer, song.Genre, song.Duration)

	if err != nil {
		return 0, exceptions.NewHTTPError(err, 400, "fail to update song")
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, exceptions.NewHTTPError(err, 400, "fail to check rows effected")
	}
	if rowsAffected == 0 {
		return 0, exceptions.NewHTTPError(err, 404, "song id not found")
	}
	return rowsAffected, nil
}

func DeleteSongService(id string) (int64, error) {
	db := database.CreateConnection()
	defer db.Close()

	query := "DELETE FROM songs WHERE id=$1"
	res, err := db.Exec(query, id)

	if err != nil {
		return 0, exceptions.NewHTTPError(err, 400, "fail to delete song")
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, exceptions.NewHTTPError(err, 400, "fail to check rows effected")
	}
	if rowsAffected == 0 {
		return 0, exceptions.NewHTTPError(err, 404, "song id not found")
	}
	return rowsAffected, nil
}

func GetSongByAlbumId(id string) ([]models.MappedSong, error) {
	db := database.CreateConnection()
	defer db.Close()

	var songs []models.MappedSong
	songs = make([]models.MappedSong, 0, 1)

	query := "SELECT id, title, performer FROM songs WHERE album_id=$1"

	rows, err := db.Query(query, id)

	if err != nil {
		return songs, err
	}
	defer rows.Close()

	for rows.Next() {
		var song models.MappedSong

		err = rows.Scan(&song.Id, &song.Title, &song.Performer)
		if err != nil {
			return songs, err
		}
		songs = append(songs, song)
	}

	return songs, err
}

func GetAllSongsByTitleService(title string) ([]models.MappedSong, error) {
	db := database.CreateConnection()
	defer db.Close()

	var songs []models.MappedSong
	songs = make([]models.MappedSong, 0, 1)

	query := "SELECT id, title, performer FROM songs WHERE LOWER(title) LIKE '%' || $1 || '%'"

	rows, err := db.Query(query, strings.ToLower(title))

	if err != nil {
		return songs, err
	}
	defer rows.Close()

	for rows.Next() {
		var song models.MappedSong

		err = rows.Scan(&song.Id, &song.Title, &song.Performer)
		if err != nil {
			return songs, err
		}
		songs = append(songs, song)
	}

	return songs, err
}

func GetAllSongsByPerformerService(performer string) ([]models.MappedSong, error) {
	db := database.CreateConnection()
	defer db.Close()

	var songs []models.MappedSong
	songs = make([]models.MappedSong, 0, 1)

	query := "SELECT id, title, performer FROM songs WHERE LOWER(performer) LIKE '%' || $1 || '%'"

	rows, err := db.Query(query, strings.ToLower(performer))

	if err != nil {
		return songs, err
	}
	defer rows.Close()

	for rows.Next() {
		var song models.MappedSong

		err = rows.Scan(&song.Id, &song.Title, &song.Performer)
		if err != nil {
			return songs, err
		}
		songs = append(songs, song)
	}

	return songs, err
}

func GetAllSongsByTitleAndPerformerService(title string, performer string) ([]models.MappedSong, error) {
	db := database.CreateConnection()
	defer db.Close()

	var songs []models.MappedSong
	songs = make([]models.MappedSong, 0, 1)

	query := "SELECT id, title, performer FROM songs WHERE LOWER(title) LIKE '%' || $1 || '%' AND LOWER(performer) LIKE '%' || $2 || '%'"

	rows, err := db.Query(query, strings.ToLower(title), strings.ToLower(performer))

	if err != nil {
		return songs, err
	}
	defer rows.Close()

	for rows.Next() {
		var song models.MappedSong

		err = rows.Scan(&song.Id, &song.Title, &song.Performer)
		if err != nil {
			return songs, err
		}
		songs = append(songs, song)
	}

	return songs, err
}
