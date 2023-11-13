package dbutils

import (
	"context"
	"database/sql"
	"log"
	"os"

	"Matthieu-OD/card_game_sixty_six/cmd/game"
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

	stmt, err := db.PrepareContext(ctx, string(sql))
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, gameid)
	return err
}

func GetOpponentReady(db *sql.DB, ctx context.Context, gameid string) bool {
	sql, err := os.ReadFile(getPathSQL() + "get_opponent_ready.sql")
	if err != nil {
		log.Fatal(err)
	}

	var opponentReady bool
	err = db.QueryRowContext(ctx, string(sql), gameid).Scan(&opponentReady)
	return opponentReady
}

func UpdateOpponentReady(db *sql.DB, ctx context.Context, gameid string, opponentReady bool) error {
	sql, err := os.ReadFile(getPathSQL() + "update_opponent_ready.sql")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.PrepareContext(ctx, string(sql))
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, opponentReady, gameid)

	return err
}

func GetGame(db *sql.DB, ctx context.Context, gameid string) game.Game {
	sql, err := os.ReadFile(getPathSQL() + "get_game.sql")
	if err != nil {
		log.Fatal(err)
	}

	var game game.Game
	err = db.QueryRowContext(ctx, string(sql), gameid).Scan(
		&game.GameID,
		&game.OpponentReady,
		&game.Asset,
		&game.RoundScore1,
		&game.RoundScore2,
		&game.TotalScore1,
		&game.TotalScore2,
		&game.Turn,
	)

	if err != nil {
		log.Fatal(err)
	}

	return game
}
