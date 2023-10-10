CREATE TABLE {{.TableName}} (
	game_id UUID NOT NULL,
	card_id INT NOT NULL,
	FOREIGN KEY (game_id) REFERENCES game(id),
	FOREIGN KEY (card_id) REFERENCES card(id),
	PRIMARY KEY (game_id, card_id)
)
