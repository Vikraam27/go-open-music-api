package services

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/Vikraam27/go-open-music-api/database"
	"github.com/Vikraam27/go-open-music-api/models"
)

func AddSongService(song models.SongPayload) (string, error) {
	db := database.CreateConnection()
	defer db.Close()

	id := fmt.Sprintf("song-%v", strconv.Itoa(rand.Intn(9999999)))
	query := "INSERT INTO songs (id, title, year, performer, genre, duration) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	var songId string
	err := db.QueryRow(query, id, song.Title, song.Year, song.Performer, song.Genre, song.Duration).Scan(&songId)

	if err != nil {
		fmt.Print(err)
		return "", fmt.Errorf("fail to add song : %v", err)
	}
	return id, nil
}
