package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func getPathSQL() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path + "/scripts/sql/game/"
}

func CreateEmptyGame(db *sql.DB, ctx context.Context, gameid string) error {
	sql, err := os.ReadFile(getPathSQL() + "create_new_empty_game.sql")
	if err != nil {
		log.Fatal(err)
	}
	res, err := db.ExecContext(ctx, string(sql), gameid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return err
}

func GetGame(db *sql.DB, ctx context.Context, gameid string) *sql.Row {
	sql, err := os.ReadFile(getPathSQL() + "get_game.sql")
	if err != nil {
		log.Fatal(err)
	}
	row := db.QueryRowContext(ctx, string(sql), gameid)
	fmt.Println(row)
	return row
}
