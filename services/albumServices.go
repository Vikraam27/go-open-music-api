package services

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Vikraam27/go-open-music-api/database"
	"github.com/Vikraam27/go-open-music-api/exceptions"
	"github.com/Vikraam27/go-open-music-api/models"
)

func CreateAlbumService(album models.AlbumPayload) (string, error) {
	db := database.CreateConnection()
	defer db.Close()
	rand.Seed(time.Now().UnixNano())

	id := fmt.Sprintf("album-%v", strconv.Itoa(rand.Intn(9999999)))
	query := "INSERT INTO albums VALUES ($1, $2, $3) RETURNING id"
	var albumId string
	err := db.QueryRow(query, id, album.Name, album.Year).Scan(&albumId)

	if err != nil {
		fmt.Print(err)
		return "", fmt.Errorf("fail to create album : %v", err)
	}
	return id, nil
}

func GetAlbumDetailService(id string) (models.Album, error) {
	db := database.CreateConnection()
	defer db.Close()

	var album models.Album

	query := "SELECT * FROM albums WHERE id = $1"
	row := db.QueryRow(query, id)
	err := row.Scan(&album.ID, &album.Name, &album.Year)

	switch err {
	case sql.ErrNoRows:
		return album, exceptions.NewHTTPError(err, 404, "album not found")
	case nil:
		return album, nil
	default:
		return album, err
	}
}

func UpdateAlbumService(id string, album models.AlbumPayload) error {
	db := database.CreateConnection()
	defer db.Close()

	query := `UPDATE albums SET name=$2, year=$3 WHERE id=$1`

	_, err := db.Exec(query, id, album.Name, album.Year)

	switch err {
	case sql.ErrNoRows:
		return exceptions.NewHTTPError(err, 404, "album not found")
	case nil:
		return nil
	default:
		return exceptions.NewHTTPError(err, 400, "fail to update album")
	}
}