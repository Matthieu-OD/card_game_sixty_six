CREATE TABLE IF NOT EXISTS player_cards (
	player_id INT CHECK(player_id IN (1, 2)) NOT NULL,
	game_id TEXT NOT NULL,
	card_id INT NOT NULL,
	FOREIGN KEY (game_id) REFERENCES game(id),
	FOREIGN KEY (card_id) REFERENCES card(id),
	PRIMARY KEY (game_id, card_id)
)
