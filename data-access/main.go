package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, os.Getenv("DBUSER"), os.Getenv("DBPASS"), "postgres")
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	album, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", album)

	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  84.22,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

}

// albumsBy Artist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM recordings.album WHERE artist = $1", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	//Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM recordings.album WHERE id = $1", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsByID %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	var lastInsertId int64
	err := db.QueryRow("INSERT INTO recordings.album (title, artist, price) VALUES ($1, $2, $3) RETURNING id",
		alb.Title, alb.Artist, alb.Price).Scan(&lastInsertId)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return lastInsertId, nil
}
