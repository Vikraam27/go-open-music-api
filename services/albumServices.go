package services

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Vikraam27/go-open-music-api/database"
	"github.com/Vikraam27/go-open-music-api/exceptions"
	"github.com/Vikraam27/go-open-music-api/models"
)

func CreateAlbumService(album models.AlbumPayload) (string, error) {
	db := database.CreateConnection()
	defer db.Close()

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

func UpdateAlbumService(id string, album models.AlbumPayload) (int64, error) {
	db := database.CreateConnection()
	defer db.Close()

	query := `UPDATE albums SET name=$2, year=$3 WHERE id=$1`

	res, err := db.Exec(query, id, album.Name, album.Year)

	if err != nil {
		return 0, exceptions.NewHTTPError(err, 400, "fail to update album")
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, exceptions.NewHTTPError(err, 400, "fail to check rows effected")
	}
	if rowsAffected == 0 {
		return 0, exceptions.NewHTTPError(err, 404, "album id not found")
	}
	return rowsAffected, nil
}

func DeleteAlbumService(id string) (int64, error) {
	db := database.CreateConnection()
	defer db.Close()

	query := `DELETE FROM albums WHERE id=$1`

	res, err := db.Exec(query, id)

	if err != nil {
		return 0, exceptions.NewHTTPError(err, 400, "fail to delete album")
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, exceptions.NewHTTPError(err, 400, "fail to check rows effected")
	}
	if rowsAffected == 0 {
		return 0, exceptions.NewHTTPError(err, 404, "album id not found")
	}
	return rowsAffected, nil
}
