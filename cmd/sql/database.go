package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strings"

	_ "modernc.org/sqlite"

	"Matthieu-OD/card_game_sixty_six/server/game"
)

// TODO: create add, get, update, delete functions
func CreateDB() *sql.DB {
	// NOTE: use in memory database in production?
	db, err := sql.Open("sqlite", "./server/sql/database.db?_foreign_keys=(1)")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to the sqlite database!")
	}
	return db
}

func CreateTables(ctx context.Context, db *sql.DB) {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	sqlFilesPath := path + "/server/sql/tables/"

	// TODO: use a global variable for the path to sql files
	files, err := os.ReadDir(sqlFilesPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		data, err := os.ReadFile(sqlFilesPath + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		// TODO: take care of the template files
		if strings.Contains(file.Name(), "template") {
			continue
		}

		_, err = db.ExecContext(ctx, string(data))
		if err != nil {
			log.Println("file name: ")
			log.Fatal("filename:", file.Name(), "error: ", err)
		}
	}
}

func PopulateDB(ctx context.Context, db *sql.DB) {
	// check if the card are already inserted
	var count int
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM card").Scan(&count)

	if count >= 28 {
		return
	}

	cardsData, err := os.ReadFile("./server/sql/data/cards.json")
	if err != nil {
		log.Fatal(err)
	}

	cardsJson := []game.Card{}

	err = json.Unmarshal(cardsData, &cardsJson)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO card (id, suit, value, points) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, card := range cardsJson {
		_, err := stmt.ExecContext(ctx, card.ID, card.Suit, card.Value, card.Points)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Populated the database with 28 cards\n")
}

func SetupDB() (*sql.DB, context.Context) {
	ctx := context.Background()

	db := CreateDB()
	CreateTables(ctx, db)
	PopulateDB(ctx, db)
	return db, ctx
}
