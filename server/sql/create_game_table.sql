CREATE TABLE game (
	id UUID NOT NULL PRIMARY KEY,
	opponent_ready BOOLEAN NOT NULL DEFAULT FALSE,

	asset_card_id INT,
	FOREIGN KEY (asset_card_id) REFERENCES card(id),

	round_score1 INT NOT NULL DEFAULT 0,
	round_score2 INT NOT NULL DEFAULT 0,
	total_score1 INT NOT NULL DEFAULT 0,
	total_score2 INT NOT NULL DEFAULT 0,

	OrderStatus TEXT CHECK(OrderStatus IN ('player1', 'player2')) NOT NULL DEFAULT 'player1',
)
