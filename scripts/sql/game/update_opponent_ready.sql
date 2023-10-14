UPDATE game SET opponent_ready = TRUE WHERE id = ?;


CREATE TABLE IF NOT EXISTS game (
	id TEXT NOT NULL PRIMARY KEY,
	opponent_ready BOOLEAN NOT NULL DEFAULT FALSE,

	asset_card_id INTEGER,
	round_score1 INTEGER NOT NULL DEFAULT 0,
	round_score2 INTEGER NOT NULL DEFAULT 0,
	total_score1 INTEGER NOT NULL DEFAULT 0,
	total_score2 INTEGER NOT NULL DEFAULT 0,

	turn TEXT CHECK(turn IN ('player1', 'player2')) NOT NULL DEFAULT 'player1',

	FOREIGN KEY (asset_card_id) REFERENCES card(id)
)
