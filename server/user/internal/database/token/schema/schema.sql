CREATE TABLE IF NOT EXISTS Tokens (
	state TEXT PRIMARY KEY,
	userId TEXT NOT NULL,
	expires TEXT NOT NULL
);
