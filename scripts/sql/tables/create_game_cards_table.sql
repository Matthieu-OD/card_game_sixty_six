CREATE TABLE IF NOT EXISTS game_cards (
	card_type VARCHAR(15) CHECK(card_type IN ('playing', 'stack', 'last_fold')) NOT NULL,
	game_id TEXT NOT NULL,
	card_id INT NOT NULL,
	FOREIGN KEY (game_id) REFERENCES game(id),
	FOREIGN KEY (card_id) REFERENCES card(id),
	PRIMARY KEY (game_id, card_id)
)
